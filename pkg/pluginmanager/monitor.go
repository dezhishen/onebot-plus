package pluginmanager

import (
	"context"
	"time"

	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/sirupsen/logrus"
)

func startMonitor(ctx context.Context) {
	go monitorDir(getRootDir(), ctx)
}

func monitorDir(path string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logrus.Infof("ctx done,exits monitor [%v]", path)
			return
		default:
			checkAllByPath(path)
		}
		time.Sleep(30 * time.Second)
	}
}

func checkAllByPath(path string) {
	scanPath(path, func(filePath string) error {
		p := GetPluginByFilePath(filePath)
		if p == nil {
			//register new?
			p, rpcCli, err := plugin.LoadPlugin(filePath)
			if err != nil {
				logrus.Warnf("register new plugin [%v] err ,msg:[%v]", filePath, err)
				return nil
			}
			Register(filePath, p, rpcCli, &cli.OnebotCliRealImpl{})
		}
		return nil
	})
	plugins := GetAllPlugins()
	for _, p := range plugins {
		c, err := p.RPClient.Client()
		if err != nil {
			p.Status = PluginStatusUnhealthy
			continue
		}
		err = c.Ping()
		if err != nil {
			p.Status = PluginStatusUnhealthy
			continue
		}
		p.Status = PluginStatusHealthy
	}
}
