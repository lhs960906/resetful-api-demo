package conf_test

import (
	"os"
	"testing"

	"github.com/lhs960906/restful-api-demo/conf"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)

	err := conf.LoadConfigFromToml("../etc/demo.toml")

	if should.NoError(err) {
		should.Equal("go_course", conf.C().MySQL.UserName)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)

	os.Setenv("MYSQL_HOST", "127.0.0.1")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_DATABASE", "unit_test")
	os.Setenv("MYSQL_USERNAME", "lhs")
	os.Setenv("MYSQL_PASSWORD", "lhs960906")

	err := conf.LoadConfigFromEnv()

	if should.NoError(err) {
		should.Equal("127.0.0.1", conf.C().MySQL.Host)
		should.Equal("3306", conf.C().MySQL.Port)
		should.Equal("unit_test", conf.C().MySQL.Database)
		should.Equal("lhs", conf.C().MySQL.UserName)
		should.Equal("lhs960906", conf.C().MySQL.Password)
	}
}

// 测试从配置文件中加载配置并初始化数据库客户端
func TestGetDB(t *testing.T) {
	should := assert.New(t)
	// os.Setenv("MYSQL_HOST", "192.168.111.50")
	// os.Setenv("MYSQL_PORT", "3306")
	// os.Setenv("MYSQL_USERNAME", "root")
	// os.Setenv("MYSQL_PASSWORD", "123456")
	// os.Setenv("MYSQL_DATABASE", "test")

	// err := conf.LoadConfigFromEnv()
	err := conf.LoadConfigFromToml("../etc/demo.toml")

	if should.NoError(err) {
		conf.C().MySQL.GetDB()
	}
}
