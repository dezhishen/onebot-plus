package main

import (
	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/sirupsen/logrus"
)

func main() {
	plugin.OnebotPluginBuilder().
		//设置插件内容
		Id("test").
		Name("测试插件").
		Description("这是测试插件").
		Help("测试查询,无操作指令").
		Init(func(cli api.OnebotApiClientInterface) error {
			result, err := cli.GetLoginInfo()
			if err != nil {
				return err
			}
			logrus.Infof(result.Data.Nickname)
			return nil
		}).
		HandleMessagePrivate(func(data *model.EventMessagePrivate, onebotApi api.OnebotApiClientInterface) error {
			onebotApi.SendPrivateMsg(
				&model.PrivateMsg{
					UserId:  data.Sender.UserId,
					Message: data.Message,
				},
			)
			return nil
		}).
		//构建插件
		Build().
		Start()
}
