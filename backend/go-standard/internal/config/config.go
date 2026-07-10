package config

type Config struct {
	Server ServerConfig
	DB     DBConfig
	Redis  RedisConfig
	Logger LogConfig
}

type ServerConfig struct {
	// 服务名称
	Name string
	// 服务端口号
	Port string
}

type LogConfig struct {
	// 日志级别，同slog的日志级别
	Level string
	// 日志显示格式化，json方式显示还是text
	Format string
}

type DBConfig struct {
	// 主机ip
	Host string
	// 端口号
	Port string
	// 用户名
	User string
	// 密码
	Pwd string
	//数据库名
	Name string
	// 数据库连接是否加密
	SSLMode string
}

type RedisConfig struct {
	// Redis 地址
	Addr string
	// Redis 密码，允许为空
	Password string
	// Redis 数据库编号
	DB string
}
