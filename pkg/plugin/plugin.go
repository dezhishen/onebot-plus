package plugin

import (
	"context"

	"github.com/dezhishen/onebot-plus/pkg/plugin/base"
	"github.com/dezhishen/onebot-plus/pkg/plugin/event"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"onebot_grpc": &OnebotGRPCPlugin{},
}

// 完整插件对象
type OnebotGRPCPlugin struct {
	// 需要嵌入插件接口
	plugin.Plugin
	// 具体实现，仅当业务接口实现基于Go时该字段有用
	Impl OnebotInterface
}

// 插件实现GRPC的接口
func (p *OnebotGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	event.RegisterOnebotEventCallbackGRPCServer(s, &event.OnebotEventCallbackServerStub{
		Broker: broker,
		Impl:   p.Impl,
	})
	base.RegisterOnebotPluginBaseGRPCServer(
		s,
		&base.OnebotPluginBaseServerStub{
			Broker: broker,
			Impl:   p.Impl,
		},
	)
	return nil
}

// 插件实现GRPC的接口
func (p *OnebotGRPCPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	baseClientSub := base.OnebotPluginBaseClientStub{
		Broker: broker,
		Client: base.NewOnebotPluginBaseGRPCClient(c),
	}
	eventClientSub := event.OnebotEventCallbackClientStub{
		Broker: broker,
		Client: event.NewOnebotEventCallbackGRPCClient(c),
	}
	result := &OnebotClientStub{
		OnebotPluginBaseClientStub:    baseClientSub,
		OnebotEventCallbackClientStub: eventClientSub,
	}
	return result, nil
}
