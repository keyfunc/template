package app

import (
	"mall/internal/domain/todo"
	"net/http"
)

// 各个领域的依赖显示传入集中管理构造
func NewRoutes(app *App) {
	// 【todo】模块
	todo.NewRoutes(&todo.TodoDeps{
		Cfg:    app.Config,
		Logger: app.Logger,
		DB:     app.DB,
		Mux:    app.Mux,
	})
	// 注册一个 openapi.json 获取接口, 用于导入接口文档
	app.Mux.HandleFunc("GET /openapi.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/openapi.json")
	})
}
