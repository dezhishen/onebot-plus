package main

import (
	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func main() {
	plugin.NewOnebotEventPluginBuilder().
		//设置插件内容
		Id("test").Name("测试插件").Description("这是测试插件").
		Help("测试查询,无操作指令").
		Init(func(cli cli.OnebotCli) error {
			sendGroupForwardMessage(cli)
			return nil
		}).
		//构建插件
		Build().
		//启动
		Start()
}

func sendGroupForwardMessage(cli cli.OnebotCli) {
	r, e := cli.GetLoginInfo()
	if e != nil {
		println(e)
		return
	}
	cli.SendGroupForwardMessageByRawMsg(398415189, r.Data.UserId, r.Data.Nickname, []*model.MessageSegment{
		{
			Type: "text",
			Data: &model.MessageElementText{
				Text: "hello,我是合并消息",
			},
		},
	})
}
