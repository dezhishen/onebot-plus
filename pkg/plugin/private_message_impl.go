package plugin

import (
	"log"

	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
)

//插件实现
type PrivateMessageRealImpl struct {
	Callback func(*model.EventMessagePrivate) error
}

func (p *PrivateMessageRealImpl) Handle(msg *model.EventMessagePrivate) error {
	return p.Callback(msg)
}

//生成插件
func NewPrivateMessageEventHandlerPlugin(callback func(msg *model.EventMessagePrivate) error) *PrivateMessageRealImpl {
	result := &PrivateMessageRealImpl{
		Callback: callback,
	}
	return result
}

//启动插件
func (p *PrivateMessageRealImpl) Start() {
	log.Print("start...")
	var pluginMap = map[string]plugin.Plugin{
		"main": &PrivateMessageGRPCPlugin{Impl: p},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
	log.Print("exit...")
}
