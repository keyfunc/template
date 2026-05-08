package response

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type PageData[T any] struct {
	List  []T   `json:"list"`
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Total int64 `json:"total"`
}

// 统一的JSON响应函数
func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)               // 设置 http状态
	_ = json.NewEncoder(w).Encode(data) // 转json 返回给客户端
}

// 成功的响应
func Success[T any](w http.ResponseWriter, data T) {
	JSON(w, http.StatusOK, Response[T]{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// 失败的响应
func Fail(w http.ResponseWriter, code int, message string) {
	JSON(w, http.StatusOK, Response[any]{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func Page[T any](w http.ResponseWriter, list []T, page int, size int, total int64) {
	Success(w, PageData[T]{
		List:  list,
		Page:  page,
		Size:  size,
		Total: total,
	})
}
