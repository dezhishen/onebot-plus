package main

import (
	"github.com/hashicorp/go-plugin"

	onebotplugin "github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("start...")
	p := onebotplugin.NewPrivateMessageEventHandlerPlugin(func(msg *model.EventMessagePrivate) error {
		logrus.Infof("收到来自%v的消息:%v", msg.Sender.Nickname, msg.Message)
		return nil
	})
	var pluginMap = map[string]plugin.Plugin{
		"main": &onebotplugin.PrivateMessageGRPCPlugin{Impl: p},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: onebotplugin.HandshakeConfig,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}
