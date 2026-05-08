package request

import (
	"fmt"
	"net/http"
	"strconv"
)

// 校验分页请求参数
func CheckPageQuery(r *http.Request) (PageQuery, error) {
	q := r.URL.Query()

	page := 1
	size := 10

	if v := q.Get("page"); v != "" {
		d, err := strconv.Atoi(v)
		// string => int 出错
		if err != nil {
			return PageQuery{}, fmt.Errorf("你输入的%d是非数值类型", d)
		}
		// 数值越界判断
		if d <= 0 {
			return PageQuery{}, fmt.Errorf("page 必须大于0")
		}
		page = d
	}
	if v := q.Get("size"); v != "" {
		d, err := strconv.Atoi(v)
		// string => int 出错
		if err != nil {
			return PageQuery{}, fmt.Errorf("你输入的%d是非数值类型", d)
		}
		// 数值越界判断
		if d <= 0 {
			return PageQuery{}, fmt.Errorf("size 必须大于0")
		}
		size = d
	}
	return PageQuery{
		Page: page,
		Size: size,
	}, nil
}
