package pluginmanager

import (
	"fmt"
	"strings"

	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-plus/pkg/command"
	"github.com/dezhishen/onebot-sdk/pkg/event"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type manageCommand struct {
	Id string `long:"id" description:"ID" default:"-1"`
}

func registerManageEventToWebsocket(onebotCli cli.OnebotCli) error {
	event.ListenMessageGroup(func(data model.EventMessageGroup) error {
		if data.Sender.Role == "admin" {
			ok, msgText := genMsgText(data.Message)
			if !ok {
				return nil
			}
			onebotCli.SendGroupMsg(genGroupTextMsg(data.GroupId, msgText))
		}
		return nil
	})
	event.ListenMessagePrivate(func(data model.EventMessagePrivate) error {
		ok, msgText := genMsgText(data.Message)
		if !ok {
			return nil
		}
		onebotCli.SendPrivateMsg(genPrivateTextMsg(data.Sender.UserId, msgText))
		return nil
	})
	return nil
}

func genPrivateTextMsg(userId int64, text string) *model.PrivateMsg {
	return &model.PrivateMsg{
		UserId: userId,
		Message: []*model.MessageSegment{
			{
				Type: "text",
				Data: &model.MessageElementText{
					Text: text,
				},
			},
		},
	}
}

func genGroupTextMsg(groupId int64, text string) *model.GroupMsg {
	return &model.GroupMsg{
		GroupId: groupId,
		Message: []*model.MessageSegment{
			{
				Type: "text",
				Data: &model.MessageElementText{
					Text: text,
				},
			},
		},
	}
}

func genMsgText(messgae []*model.MessageSegment) (bool, string) {
	if len(messgae) == 0 {
		return false, ""
	}
	if messgae[0].Type != "text" {
		return false, ""
	}
	v, ok := messgae[0].Data.(*model.MessageElementText)
	if !ok {
		return false, ""
	}
	if !strings.HasPrefix(v.Text, ".onebot") {
		return false, ""
	}
	commandArgs := manageCommand{}
	commands, err := command.Parse(".onebot", &commandArgs, strings.Split(v.Text, " "))
	if err != nil {
		return false, fmt.Sprintf("%v", err)
	}
	for i, command := range commands {
		if i == 0 {
			continue
		}
		if command == "list" {
			plugins := GetAllPlugins()
			if len(plugins) == 0 {
				return true, "当前没有启用的插件"
			}
			var msgText string
			for i, p := range plugins {
				msgText += fmt.Sprintf("[%v] %v %v\n", i+1, p.Plugin.Id(), p.Plugin.Name())
			}
			msgText += "请输入[.onebot help --id $id]查看指定插件的帮助指令"
			return true, msgText
		}
		if command == "help" {
			if commandArgs.Id == "" {
				return true, "请输入需要查看帮助的插件ID"
			}
			p := GetPluginById(commandArgs.Id)
			if p == nil {
				return true, fmt.Sprintf("要操作的插件[%v]不存在", commandArgs.Id)
			}
			if p.Plugin.Help() == "" {
				return true, fmt.Sprintf("[%v]:%v", p.Plugin.Name(), "没有发现帮助说明")
			}
			return true, fmt.Sprintf("[%v]:%v", p.Plugin.Name(), p.Plugin.Help())
		}
	}
	return false, ""
}
