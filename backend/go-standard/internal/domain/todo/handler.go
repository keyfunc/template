package todo

import (
	"encoding/json"
	"log/slog"
	"mall/internal/pkg/request"
	"mall/internal/pkg/response"
	"net/http"
	"strconv"
)

type Handler struct {
	Logger  *slog.Logger
	Service *Service
}
type HandlerDeps struct {
	Logger  *slog.Logger
	Service *Service
}

func NewHandler(deps HandlerDeps) *Handler {
	return &Handler{
		Logger:  deps.Logger,
		Service: deps.Service,
	}
}

// todo 查询
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	// 1. 分页请求参数校验
	var pageQuery request.PageQuery
	pageQuery, err := request.CheckPageQuery(r)
	if err != nil {
		h.Logger.Error("参数错误", "error", err)
		response.Fail(w, http.StatusBadRequest, err.Error())
		return
	}
	// 2.调用 service 层的关于查询todo的业务逻辑
	res, err := h.Service.List(r.Context(), ListTodoReq{
		Page: pageQuery.Page,
		Size: pageQuery.Size,
	})
	if err != nil {
		h.Logger.Error("查询错误", "error", err)
		response.Fail(w, http.StatusBadRequest, err.Error())
		return
	}
	// 3. 成功的话返回查询数据
	response.Page(w, res.List, pageQuery.Page, pageQuery.Size, res.Total)
}

// todo 新建
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	// 1. 参数校验， post参数从body中取值
	var req CreateTodoReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Fail(w, http.StatusBadRequest, err.Error())
		return
	}
	if req.Title == "" {
		response.Fail(w, http.StatusBadRequest, "标题不能为空")
		return
	}
	if req.Description == "" {
		response.Fail(w, http.StatusBadRequest, "描述不能为空")
		return
	}
	// 2. 调用service层的创建todo的业务逻辑
	todo, err := h.Service.Create(r.Context(), req)
	if err != nil {
		response.Fail(w, http.StatusBadRequest, err.Error())
	}
	response.Success(w, todo)
}

// 更新 todo
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	// 1. path id 校验
	idStr := r.PathValue("id") // 路径的id string类型，需要转换为整型
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		response.Fail(w, http.StatusBadRequest, "id 参数错误")
		return
	}

	// 2. 请求体 参数校验
	var req UpdateTodoReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Fail(w, http.StatusBadRequest, err.Error())
		return
	}
	if req.Title == "" {
		response.Fail(w, http.StatusBadRequest, "标题不能为空")
		return
	}
	if req.Description == "" {
		response.Fail(w, http.StatusBadRequest, "描述不能为空")
		return
	}
	if req.Status == 0 {
		response.Fail(w, http.StatusBadRequest, "状态不能为空")
		return
	}

	// 2. 交给service 层处理业务逻辑
	res, err1 := h.Service.Update(r.Context(), id, req)
	if err1 != nil {
		response.Fail(w, http.StatusBadRequest, err1.Error())
		return
	}
	// 3. 成功返回
	response.Success(w, res)
}

// 单个删除
func (h *Handler) Del(w http.ResponseWriter, r *http.Request) {
	// 1. id 参数校验
	idStr := r.PathValue("id") // 路径的id string类型，需要转换为整型
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		response.Fail(w, http.StatusBadRequest, "id 参数错误")
		return
	}
	// 2. 待用业务层处理业务逻辑
	err1 := h.Service.Del(r.Context(), id)
	if err1 != nil {
		response.Fail(w, http.StatusBadRequest, err1.Error())
		return
	}
	// 3. 删除成功返回
	response.Success(w, "")
}
