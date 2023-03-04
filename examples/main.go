package main

import (
	"io"
	"time"

	"github.com/dezhishen/onebot-plus/pkg/pluginmanager"
	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/config"
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
	d, c, err := pluginmanager.LoadPlugin(path)
	if err != nil {
		panic(err)
	}
	defer c.Kill()
	logrus.Infof("插件启动,id[%v],name[%v]", d.Id(), d.Name())
	d.Init(channelCli)
	time.Sleep(10 * time.Second)
	d.BeforeExit(channelCli)
}
