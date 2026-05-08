package todo

// todo 查询参数
type ListTodoReq struct {
	Page int
	Size int
}

// todo 查询响应参数
type ListTodoRes struct {
	List  []Todo `json:"list"`
	Total int64  `json:"total"`
}

// create Todo 请求参数
type CreateTodoReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// update Todo 请求参数
type UpdateTodoReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}
