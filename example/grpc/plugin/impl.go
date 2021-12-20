package main

import (
	"github.com/dezhishen/onebot-plus/example/grpc/common"
	onebotPlugin "github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
)

type DemoGRPCRelImp struct {
}

func (d DemoGRPCRelImp) SayHi(name string) (string, error) {
	logrus.Infof("recive call name : %v", name)
	return "Hi " + name + " !", nil
}

func main() {
	demo := &DemoGRPCRelImp{}
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"demo": &common.DemoGRPCPlugin{Impl: demo},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: onebotPlugin.HandshakeConfig,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}
