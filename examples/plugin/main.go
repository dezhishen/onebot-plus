package main

import (
	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/sirupsen/logrus"
)

func main() {
	path := "./impl/plugin.exe"
	d, c := plugin.LoadPlugin(path)
	defer c.Kill()
	logrus.Infof("插件启动,id[%v],name[%v]", d.Id(), d.Name())
}
