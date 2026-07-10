package logger

import (
	"log/slog"
	"mall/internal/config"
	"os"
)

func New(cfg config.LogConfig) *slog.Logger {
	var level slog.Level

	switch cfg.Level {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: level,
		// AddSource: true, // 需要打印源码位置时可以打开
	}

	var handler slog.Handler
	switch cfg.Format {
	case "json":
		// 生产环境建议使用 JSON 格式，方便日志收集和分析
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		// 开发环境建议使用文本格式，便于阅读
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	log := slog.New(handler)
	slog.SetDefault(log)

	return log
}
