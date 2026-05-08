package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Config, error) {

	data, err := os.ReadFile(path) // 读取文件内容，返回字节切片
	if err != nil {
		return nil, err
	}
	var cfg Config
	// 反序列化，将yaml文件数据转为go数据
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
