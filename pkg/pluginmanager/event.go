package pluginmanager

import (
	"fmt"
	"strings"

	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-plus/pkg/command"
	"github.com/dezhishen/onebot-plus/pkg/env"
	"github.com/dezhishen/onebot-sdk/pkg/event"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type manageCommand struct {
	Id string `long:"id" description:"ID" default:"-1"`
}

func registerPluginEventToWebsocket() {
	event.ListenMessageGroup(func(data model.EventMessageGroup) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.MessageGroup(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenMessagePrivate(func(data model.EventMessagePrivate) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.MessagePrivate(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenMetaHeartBeat(func(data model.EventMetaHeartbeat) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.MetaHeartbeat(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenMetaLifecycle(func(data model.EventMetaLifecycle) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.MetaLifecycle(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenNoticeGroupUpload(func(data model.EventNoticeGroupUpload) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.NoticeGroupUpload(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenNoticeFriendAdd(func(data model.EventNoticeFriendAdd) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.NoticeFriendAdd(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenNoticeFriendRecall(func(data model.EventNoticeFriendRecall) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.NoticeFriendRecall(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenNoticeGroupAdmin(func(data model.EventNoticeGroupAdmin) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.NoticeGroupAdmin(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenNoticeGroupBan(func(data model.EventNoticeGroupBan) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.NoticeGroupBan(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenNoticeGroupDecrease(func(data model.EventNoticeGroupDecrease) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.NoticeGroupDecrease(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenNoticeGroupIncrease(func(data model.EventNoticeGroupIncrease) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.NoticeGroupIncrease(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenNoticeGroupNotifyHonor(func(data model.EventNoticeGroupNotifyHonor) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.NoticeGroupNotifyHonor(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenNoticeGroupNotifyLuckyKing(func(data model.EventNoticeGroupNotifyLuckyKing) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.NoticeGroupNotifyLuckyKing(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenNoticeGroupNotifyPoke(func(data model.EventNoticeGroupNotifyPoke) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.NoticeGroupNotifyPoke(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenNoticeGroupRecall(func(data model.EventNoticeGroupRecall) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.NoticeGroupRecall(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenRequestFriend(func(data model.EventRequestFriend) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.RequestFriend(&data, p.OnebotCli)
		}
		return nil
	})
	event.ListenRequestGroup(func(data model.EventRequestGroup) error {
		for _, p := range GetAllPlugins() {
			if p.Status != PluginStatusHealthy {
				continue
			}
			go p.Plugin.RequestGroup(&data, p.OnebotCli)
		}
		return nil
	})
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
	if len(commands) == 0 {
		return false, ""
	}
	if len(commands) > 1 {
		command := commands[1]
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
		if command == "env" {
			if len(commands) <= 2 {
				return false, ""
			}
			var envKey string
			var envVal string
			envKey = commands[2]
			if len(command) > 3 {
				envVal = commands[3]
			}
			env.SetEnv(envKey, envVal)
		}
	}
	return false, ""
}
