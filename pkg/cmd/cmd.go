package cmd

import (
	"context"

	"github.com/dezhishen/onebot-plus/pkg/cli"
	myplugin "github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-plus/pkg/pluginmanager"
	"github.com/dezhishen/onebot-sdk/pkg/event"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
)

type pluginWithClient struct {
	plugin myplugin.OnebotEventPlugin
	client plugin.Client
}

var all_plugins = make(map[string]*pluginWithClient)

func StartWithContext(ctx context.Context) error {
	onebotCli := &cli.OnebotCliRealImpl{}
	defer closeAll(onebotCli)
	err := pluginmanager.ScanInPath("./plugins", func(file string) error {
		err := setPluginEvent(file, onebotCli)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	//开启监听
	err = event.StartWsWithContext(ctx)
	if err != nil {
		panic(err)
	}
	return nil
}

//启动
func Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	onebotCli := &cli.OnebotCliRealImpl{}
	defer closeAllWithCancel(cancel, onebotCli)
	err := pluginmanager.ScanInPath("./plugins", func(file string) error {
		err := setPluginEvent(file, onebotCli)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	//开启监听
	err = event.StartWsWithContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func closeAll(onebotCli cli.OnebotCli) {
	for _, v := range all_plugins {
		v.plugin.BeforeExit(onebotCli)
		v.client.Kill()
	}
}

func closeAllWithCancel(cancel context.CancelFunc, onebotCli cli.OnebotCli) {
	for _, v := range all_plugins {
		go closePlugin(v, onebotCli)
	}
	cancel()
}

func closePlugin(v *pluginWithClient, onebotCli cli.OnebotCli) {

	v.plugin.BeforeExit(onebotCli)
	v.client.Kill()

}

func setPluginEvent(file string, onebotCli cli.OnebotCli) error {
	p, client := myplugin.LoadPlugin(file)
	all_plugins[p.Id()] = &pluginWithClient{
		plugin: p,
		client: *client,
	}
	event.ListenMessageGroup(func(data model.EventMessageGroup) error {
		return p.MessageGroup(&data, onebotCli)
	})
	event.ListenMessagePrivate(func(data model.EventMessagePrivate) error {
		return p.MessagePrivate(&data, onebotCli)
	})
	event.ListenMetaHeartBeat(func(data model.EventMetaHeartbeat) error {
		return p.MetaHeartbeat(&data, onebotCli)
	})
	event.ListenMetaLifecycle(func(data model.EventMetaLifecycle) error {
		return p.MetaLifecycle(&data, onebotCli)
	})
	event.ListenNoticeFriendAdd(func(data model.EventNoticeFriendAdd) error {
		return p.NoticeFriendAdd(&data, onebotCli)
	})
	event.ListenNoticeFriendRecall(func(data model.EventNoticeFriendRecall) error {
		return p.NoticeFriendRecall(&data, onebotCli)
	})
	event.ListenNoticeGroupAdmin(func(data model.EventNoticeGroupAdmin) error {
		return p.NoticeGroupAdmin(&data, onebotCli)
	})
	event.ListenNoticeGroupBan(func(data model.EventNoticeGroupBan) error {
		return p.NoticeGroupBan(&data, onebotCli)
	})
	event.ListenNoticeGroupDecrease(func(data model.EventNoticeGroupDecrease) error {
		return p.NoticeGroupDecrease(&data, onebotCli)
	})
	event.ListenNoticeGroupIncrease(func(data model.EventNoticeGroupIncrease) error {
		return p.NoticeGroupIncrease(&data, onebotCli)
	})
	event.ListenNoticeGroupNotifyHonor(func(data model.EventNoticeGroupNotifyHonor) error {
		return p.NoticeGroupNotifyHonor(&data, onebotCli)
	})
	event.ListenNoticeGroupNotifyLuckyKing(func(data model.EventNoticeGroupNotifyLuckyKing) error {
		return p.NoticeGroupNotifyLuckyKing(&data, onebotCli)
	})
	event.ListenNoticeGroupNotifyPoke(func(data model.EventNoticeGroupNotifyPoke) error {
		return p.NoticeGroupNotifyPoke(&data, onebotCli)
	})
	event.ListenNoticeGroupRecall(func(data model.EventNoticeGroupRecall) error {
		return p.NoticeGroupRecall(&data, onebotCli)
	})
	event.ListenRequestFriend(func(data model.EventRequestFriend) error {
		return p.RequestFriend(&data, onebotCli)
	})
	event.ListenRequestGroup(func(data model.EventRequestGroup) error {
		return p.RequestGroup(&data, onebotCli)
	})
	return nil
}
