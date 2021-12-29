package main

import (
	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/sirupsen/logrus"
)

func main() {
	plugin.NewOnebotEventPluginBuilder().
		//设置插件内容
		Id("test").Name("测试插件").Description("这是测试插件").
		//回调事件
		MessageGroup(func(req *model.EventMessageGroup, cli cli.OnebotCli) error {
			logrus.Infof("收到群组消息,UserId:[%v]", req.Sender.UserId)
			return nil
		}).
		MessagePrivate(func(req *model.EventMessagePrivate, cli cli.OnebotCli) error {
			logrus.Infof("收到私聊消息,UserId:[%v]", req.Sender.UserId)
			return nil
		}).
		MetaHeartbeat(func(req *model.EventMetaHeartbeat, cli cli.OnebotCli) error {
			logrus.Infof("收到EventMetaHeartbeat,SelfId:[%v]", req.SelfId)
			return nil

		}).
		//构建插件
		Build().
		//启动
		Start()
}
