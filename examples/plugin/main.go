package main

import (
	"io/ioutil"
	"time"

	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetOutput(ioutil.Discard)
	path := "./impl/plugin.exe"
	d, c := plugin.LoadPlugin(path)
	defer c.Kill()
	logrus.Infof("插件启动,id[%v],name[%v]", d.Id(), d.Name())

	gMsg := &model.EventMessageGroup{}
	gMsg.Sender = &model.Sender{
		UserId: 123456,
	}
	cli := &cli.OnebotCliRealImpl{}
	d.MessageGroup(gMsg, cli)
	pM := &model.EventMessagePrivate{}
	pM.Sender = &model.Sender{
		UserId: 123456,
	}
	d.MessagePrivate(pM, cli)
	d.MetaHeartbeat(&model.EventMetaHeartbeat{}, cli)
	d.MetaLifecycle(&model.EventMetaLifecycle{}, cli)
	d.Init(cli)
	time.Sleep(100 * time.Second)
}
