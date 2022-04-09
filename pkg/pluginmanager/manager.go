package pluginmanager

import (
	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-plus/pkg/plugin"

	rpcPlugin "github.com/hashicorp/go-plugin"
)

var allPlugins = make(map[string]*containerOfPlugin)

type PluginStatus int

const (
	PluginStatusHealthy PluginStatus = iota
	PluginStatusError
)

type containerOfPlugin struct {
	Plugin    plugin.OnebotEventPlugin
	Status    PluginStatus
	RPClient  *rpcPlugin.Client
	OnebotCli cli.OnebotCli
}

func Register(plugin plugin.OnebotEventPlugin, rpcCli *rpcPlugin.Client, onebotCli cli.OnebotCli) error {
	err := plugin.Init(onebotCli)
	if err != nil {
		return err
	}
	value := &containerOfPlugin{
		Plugin:    plugin,
		Status:    PluginStatusHealthy,
		RPClient:  rpcCli,
		OnebotCli: onebotCli,
	}
	allPlugins[plugin.Id()] = value
	return nil
}

func GetPluginById(id string) *containerOfPlugin {
	if plugin, ok := allPlugins[id]; ok {
		return plugin
	}
	return nil
}

func GetAllPlugins() []*containerOfPlugin {
	if len(allPlugins) == 0 {
		return nil
	}
	var result []*containerOfPlugin
	for _, e := range allPlugins {
		result = append(result, e)
	}
	return result
}
