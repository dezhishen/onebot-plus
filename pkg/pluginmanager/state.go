package pluginmanager

import (
	"fmt"

	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-plus/pkg/plugin"

	rpcPlugin "github.com/hashicorp/go-plugin"
)

var allPlugins = make(map[string]*containerOfPlugin)
var allPluginFileInfo = make(map[string]string)

type PluginStatus int

const (
	PluginStatusHealthy PluginStatus = iota
	PluginStatusUnhealthy
)

type containerOfPlugin struct {
	Plugin    plugin.OnebotEventPlugin
	Status    PluginStatus
	RPClient  *rpcPlugin.Client
	OnebotCli cli.OnebotCli
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

func Register(filePath string, plugin plugin.OnebotEventPlugin, rpcCli *rpcPlugin.Client, onebotCli cli.OnebotCli) error {
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
	id := plugin.Id()
	if id == "" {
		return fmt.Errorf("组件[%v]未提供ID", filePath)
	}
	allPlugins[id] = value
	allPluginFileInfo[filePath] = id
	return nil
}

func GetPluginByFilePath(filePath string) *containerOfPlugin {
	if id, ok := allPluginFileInfo[filePath]; ok {
		if plugin, ok := allPlugins[id]; ok {
			return plugin
		}
	}
	return nil
}
