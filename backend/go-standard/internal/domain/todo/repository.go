package todo

import (
	"context"
	"fmt"
	"mall/internal/pkg/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}
type RepositoryDeps struct {
	DB *pgxpool.Pool
}

func NewRepository(deps RepositoryDeps) *Repository {
	return &Repository{
		DB: deps.DB,
	}
}

// 查询 todo
func (r *Repository) QueryTodo(ctx context.Context, query ListTodoReq) (*ListTodoRes, error) {

	// 1. 查询总条数,先查询总条数,再查询分页数据
	var total int64
	if err := r.DB.QueryRow(ctx, `
	SELECT COUNT(*)
	FROM todo
`).Scan(&total); err != nil {
		return nil, err
	}
	// 总条数0，返回空数组
	if total == 0 {
		return &ListTodoRes{
			List:  []Todo{},
			Total: 0,
		}, nil
	}
	// page 越界判断，参数必要性的校验已经在handler层处理了，此处只需要添加page越界判断
	maxPage := util.ClacMaxPage(total, int64(query.Size))
	if query.Page > int(maxPage) {
		err := fmt.Errorf("page 超出范围，最大页码为%d", maxPage)
		return &ListTodoRes{
			List:  []Todo{},
			Total: total,
		}, err
	}

	// 2. page size => limit offset 换算
	limit := query.Size                     // 取多少条
	offset := (query.Page - 1) * query.Size // 跳过多少条

	rows, err := r.DB.Query(ctx, `
	SELECT id,title,status,description,create_at,update_at FROM todo
	ORDER BY id
	LIMIT $1 OFFSET $2
`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 构造数组，向数组中写入数据库读取出来的数据
	var list []Todo

	for rows.Next() {
		var item Todo
		err := rows.Scan(&item.ID, &item.Title, &item.Status, &item.Description, &item.Create_at, &item.Update_at)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	// 检查循环过程中是否发现错误
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &ListTodoRes{
		List:  list,
		Total: total,
	}, nil
}

// 创建todo
func (r *Repository) CreateTodo(ctx context.Context, body CreateTodoReq) (Todo, error) {
	var todo Todo
	err := r.DB.QueryRow(ctx, `
	INSERT INTO todo (title, description)
	VALUES ($1, $2)
	RETURNING id, title, description,status, create_at,update_at
	`, body.Title, body.Description).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Create_at, &todo.Update_at)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

// 修改todo
func (r *Repository) UpdateTodo(ctx context.Context, id int64, body UpdateTodoReq) (Todo, error) {
	var todo Todo
	err := r.DB.QueryRow(ctx, `
	UPDATE todo 
	SET 
	  title = $1,
	  description = $2,
	  status = $3,
	  update_at = NOW()
	  WHERE id = $4
	  RETURNING id, title, description, status, create_at, update_at
	`, body.Title, body.Description, body.Status, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Create_at, &todo.Update_at)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

// 删除todo
func (r *Repository) DelTodo(ctx context.Context, id int64) error {
	tag, err := r.DB.Exec(ctx, `
	DELETE FROM todo
	WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("todo 不存在")
	}
	return nil
}
