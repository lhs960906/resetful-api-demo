package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lhs960906/restful-api-demo/apps/host"
)

// 接口实现的静态检查
// 这一行是为了确保 *HostServiceImple 实现了 host.Service 接口
var _ host.Service = (*HostServiceImpl)(nil)

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// Host service 服务的 Logger
		// 是对 Zap 进行了封装让其满足 Logger 接口
		// 封装原因:
		//     1. Logger 全局示例
		//     2. Logger Level 的动态调整
		//     3. 加入日志轮转的功能
		log: zap.L().Named("Host"),
	}
}

type HostServiceImpl struct {
	log logger.Logger
}
