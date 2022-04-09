package main

import (
	"io/ioutil"
	"time"

	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetOutput(ioutil.Discard)
	path := "./examples/plugin/impl/plugin-windows-amd64.exe"
	d, c := plugin.LoadPlugin(path)
	defer c.Kill()
	logrus.Infof("插件启动,id[%v],name[%v]", d.Id(), d.Name())
	d.Init(&cli.OnebotCliRealImpl{})
	time.Sleep(10 * time.Second)
	d.BeforeExit(&cli.OnebotCliRealImpl{})
}
