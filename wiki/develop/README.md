# develop

## 插件说明
插件基于回调函数进行开发
### 配置
* Id("test") //设置ID
* Name("测试插件") //设置名称
* Description("这是测试插件") //设置描述
* Help("插件帮助") //设置帮助文本
### 插件生命周期
* Init(cli cli.Bot)error //初始化 **cli的grpc连接不会断开**
* BeforeExit(cli cli.Bot)error //插件退出前回调
### Onebot事件回调
[https://github.com/botuniverse/onebot-11/tree/master/event](https://github.com/botuniverse/onebot-11/tree/master/event)
## 插件开发
### go语言
本项目提供了一套`plugin.Buidler`用于插件开发

可以定义插件的配置、插件生命周期的回调函数、onebot机器人事件的回调函数
```
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

```

### 其他语言
* proto文件
    * [plugin.proto](../../pkg/plugin/plugin.proto)
* 插件库
    * [`hashicorp/go-plugin`](https://github.com/hashicorp/go-plugin)