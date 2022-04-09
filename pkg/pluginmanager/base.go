package pluginmanager

import (
	"context"

	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/event"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func Init(ctx context.Context) error {
	err := scanPath("./plugins", func(file string) error {
		p, rpcCli := plugin.LoadPlugin(file)
		return Register(p, rpcCli, &cli.OnebotCliRealImpl{})
	})
	if err != nil {
		return err
	}
	//注册监听事件
	registerEventToWebsocket()
	//开启监听
	err = event.StartWsWithContext(ctx)
	return err

}

func registerEventToWebsocket() {
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
