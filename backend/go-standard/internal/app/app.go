package app

import (
	"log/slog"
	"mall/internal/config"
	"mall/internal/infra/db"
	"mall/internal/infra/logger"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Config *config.Config // 配置对象
	Logger *slog.Logger   // 日志
	DB     *pgxpool.Pool  // 数据库
	Mux    *http.ServeMux // net/http 分发器
}

func New() error {
	// 1. 从环境变量读取配置
	cfg, err := config.Init()
	if err != nil {
		return err
	}

	// 2.  初始化日志
	logger := logger.New(cfg.Logger)

	// 3. 初始化数据库连接
	db, err := db.NewPostgres(cfg.DB)
	if err != nil {
		return err
	}

	// 4.net/http 分发器构造
	mux := http.NewServeMux()

	// 5. 领域依赖显式传入
	NewRoutes(&App{
		Config: cfg,
		Logger: logger,
		DB:     db,
		Mux:    mux,
	})

	// 6. 端口监听
	port := ":" + cfg.Server.Port
	slog.Info("服务启动成功")
	if err := http.ListenAndServe(port, mux); err != nil {
		return err
	}

	return nil
}
