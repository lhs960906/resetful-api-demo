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
