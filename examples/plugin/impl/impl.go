package main

import (
	"time"

	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/sirupsen/logrus"
)

func main() {
	plugin.NewOnebotEventPluginBuilder().
		//设置插件内容
		Id("test").Name("测试插件").Description("这是测试插件").
		//设置生命周期函数
		Init(func(cli cli.OnebotCli) error {
			d, e := cli.GetVersionInfo()
			logrus.Info("============ version info ==========")
			logrus.Info(d, e)
			go loop(cli)
			return nil
		}).
		BeforeExit(func(cli cli.OnebotCli) error {
			d, e := cli.GetVersionInfo()
			logrus.Info("========exit=======")
			logrus.Info(d, e)
			return nil
		}).
		//Onebot回调事件
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
		MetaLifecycle(func(req *model.EventMetaLifecycle, cli cli.OnebotCli) error {
			logrus.Info("============status==========")
			logrus.Info(cli.GetStatus())
			return nil
		}).
		//构建插件
		Build().
		//启动
		Start()
}

func loop(cli cli.OnebotCli) {
	for {
		time.Sleep(1 * time.Second)
		logrus.Info("===============alive==============")
		r, e := cli.GetStatus()
		if e != nil {
			logrus.Error(e)
			continue
		}
		if r != nil {
			logrus.Infof("online:%v,good:%v", r.Online, r.Good)
		}
	}
}
