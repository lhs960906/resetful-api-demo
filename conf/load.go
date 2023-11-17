package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 完成 toml 配置文件到 Config 对象的映射
func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()

	if _, err := toml.DecodeFile(filePath, config); err != nil {
		return fmt.Errorf("load config from file error! path: %s, %s", filePath, err)
	}

	return nil
}

// 完成环境变量到 Config 对象的映射
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()

	if err := env.Parse(config); err != nil {
		return fmt.Errorf("load config from env error!, %s", err)
	}

	return nil
}
