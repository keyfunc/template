package main

import (
	"log/slog"
	"mall/internal/app"
)

func main() {
	if err := app.New(); err != nil {
		slog.Error("应用启动失败:", "error", err)
	}
}
