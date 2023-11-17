package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 完成 toml 配置文件到 Config 对象的映射
func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()

	// 加载配置文件
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error! path: %s, %s", filePath, err)
	}

	// 加载 db 全局实例
	db, err = config.MySQL.getDBConn()
	if err != nil {
		return fmt.Errorf("load db client err! error: %s", err)
	}

	return nil
}

// 完成环境变量到 Config 对象的映射
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()

	err := env.Parse(config)
	if err != nil {
		return fmt.Errorf("load config from env error!, %s", err)
	}

	// 加载 db 全局实例
	db, err = config.MySQL.getDBConn()
	if err != nil {
		return fmt.Errorf("load db client err! error: %s", err)
	}

	return nil
}
