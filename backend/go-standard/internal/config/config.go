package config

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
	Logger LogConfig    `yaml:"logger"`
}

type ServerConfig struct {
	// 服务器地址
	Addr string `yaml:"addr"`
	// 服务端口号
	Port int `yaml:"port"`
}

type LogConfig struct {
	// 日志级别，同slog的日志级别
	Level string `yaml:"level"`
	// 日志显示格式化，json方式显示还是text
	Format string `yaml:"format"`
}

type DBConfig struct {
	// 主机ip
	Host string `yaml:"host"`
	// 端口号
	Port int `yaml:"port"`
	// 用户名
	User string `yaml:"user"`
	// 密码
	Pwd string `yaml:"pwd"`
	//数据库名
	Name string `yaml:"name"`
	// 数据库连接是否加密
	SSLMode string `yaml:"sslmode"`
}
