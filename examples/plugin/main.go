package main

import (
	"io/ioutil"

	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/sirupsen/logrus"
)

type addHelper struct{}

func (*addHelper) Sum(a, b int64) (int64, error) {
	return a + b, nil
}

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
	// gMsg.Anonymous = &model.Anonymous{}
	d.MessageGroup(gMsg)
	// msg := &model.EventMessagePrivate{}
	// msg.Sender = &model.Sender{
	// 	UserId: 123456,
	// }
	// d.MessagePrivate(msg, &cli.MessageCliRealImpl{})
}
