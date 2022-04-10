package pluginmanager

import (
	"context"

	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/event"
)

func getRootDir() string {
	return "./plugins"
}

func Init(ctx context.Context) error {
	//注册插件管理事件
	registerManageEventToWebsocket(&cli.OnebotCliRealImpl{})
	err := scanPath(
		getRootDir(),
		func(filePath string) error {
			p, rpcCli, err := plugin.LoadPlugin(filePath)
			if err != nil {
				return err
			}
			return Register(filePath, p, rpcCli, &cli.OnebotCliRealImpl{})
		})
	if err != nil {
		return err
	}
	//开启监控
	startMonitor(ctx)
	//注册监听事件
	registerPluginEventToWebsocket()
	//开启监听
	err = event.StartWsWithContext(ctx)
	return err
}
