package main

import (
	"context"
	"io"
	"time"

	"github.com/dezhishen/onebot-plus/pkg/pluginmanager"
	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/event"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetOutput(io.Discard)
	path := "./examples/plugins/plugin-windows-amd64.exe"
	// load all plugins
	conf, err := config.LoadConfig("onebot-plus")
	if err != nil {
		panic(err)
	}
	channelCli, err := api.NewOnebotApiClient(conf.Api)
	if err != nil {
		panic(err)
	}
	eventCli, err := event.NewOnebotEventCli(conf.Event)
	if err != nil {
		panic(err)
	}
	plugin, c, err := pluginmanager.LoadPlugin(path)
	if err != nil {
		panic(err)
	}
	defer c.Kill()
	eventCli.ListenMessageGroup(func(data model.EventMessageGroup) error {
		plugin.HandleMessageGroup(&data, channelCli)
		return nil
	})
	logrus.Infof("插件启动,id[%v],name[%v]", plugin.Id(), plugin.Name())
	plugin.Init(channelCli)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	defer plugin.BeforeExit(channelCli)
	eventCli.StartListenWithCtx(ctx)

}
