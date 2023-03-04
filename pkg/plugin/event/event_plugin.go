package event

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type OnebotEventGRPCPlugin struct {
	// 需要嵌入插件接口
	plugin.Plugin
	// 具体实现，仅当业务接口实现基于Go时该字段有用
	Impl OnebotEventCallBackInterface
}

// 插件实现GRPC的接口
func (p *OnebotEventGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	// svcStub := &On{Impl: p.Impl}
	svcStub := &OnebotEventCallbackServerStub{
		Impl: p.Impl,
	}
	RegisterOnebotEventCallbackGRPCServer(s, svcStub)
	return nil
}

// 插件实现GRPC的接口
func (p *OnebotEventGRPCPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &OnebotEventCallbackClientStub{Client: NewOnebotEventCallbackGRPCClient(c)}, nil
}
