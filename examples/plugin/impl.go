package main

import (
	"time"

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
			go func() {
				for {
					result, err := cli.GetLoginInfo()
					if err != nil {
						logrus.Errorf("GetLoginInfo err %v", err)
					} else {
						logrus.Infof(result.Data.Nickname)
					}
					time.Sleep(time.Second * 5)
				}
			}()
			return nil
		}).
		HandleMessageGroup(func(data *model.EventMessageGroup, onebotApi api.OnebotApiClientInterface) error {
			var msgs []*model.MessageSegment
			msgs = append(msgs, &model.MessageSegment{
				Type: "text",
				Data: &model.MessageElementText{
					Text: "你好,我复读机\n",
				},
			})
			msg := append(msgs, data.Message...)
			onebotApi.SendGroupMsg(
				&model.GroupMsg{
					GroupId: data.GroupId,
					Message: msg,
				},
			)
			return nil
		}).
		//构建插件
		Build().
		Start()
}
