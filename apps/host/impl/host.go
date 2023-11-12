package impl

import (
	"context"

	"github.com/lhs960906/restful-api-demo/apps/host"
)

// 创建主机
func (host *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	host.log.Debug("create host.")
	host.log.Debugf("create host %s.", ins.Name)
	// fmt.Println("create host")
	return nil, nil
}

// 查询主机列表(这里为了方便做扩展, 引入HostSet)
func (host *HostServiceImpl) QueryHost(ctx context.Context, query *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}

// 查询主机详情
func (host *HostServiceImpl) DescribeHost(ctx context.Context, query *host.QueryHostRequest) (*host.Host, error) {
	return nil, nil
}

// 主机更新
func (host *HostServiceImpl) UpdateHost(ctx context.Context, update *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

// 主机删除
func (host *HostServiceImpl) DeleteHost(ctx context.Context, delete *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
