package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/lhs960906/restful-api-demo/apps/host"
	"github.com/lhs960906/restful-api-demo/apps/host/impl"
)

var (
	// 确保定义对象满足改接口示例
	service host.Service
)

func NewHost() *host.Host {
	return &host.Host{
		Resource: &host.Resource{},
		Describe: &host.Describe{},
	}
}

func TestCreateHost(t *testing.T) {
	ins := NewHost()
	ins.Name = "test"
	service.CreateHost(context.Background(), ins)
}

func init() {
	// 为什么不涉及默认打印日志?
	// 因为性能问题, 默认不打开日志, 如需打开, 需全局初始化日志
	fmt.Println(zap.DevelopmentSetup())
	service = impl.NewHostServiceImpl()
}
