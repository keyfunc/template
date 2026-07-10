package db

import (
	"context"
	"fmt"
	"log/slog"
	"mall/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

// 依据传入的配置初始化 postgresSQL 数据库，返回database/sql db 连接管理对象
func NewPostgres(cfg config.DBConfig) (*pgxpool.Pool, error) {
	ctx := context.Background()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Pwd,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	slog.Info("数据库连接成功")
	return pool, nil
}
