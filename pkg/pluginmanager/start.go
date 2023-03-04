package pluginmanager

import (
	"context"
	"os"
	"os/exec"

	plus_plugin "github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/event"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
)

func getRootDir() string {
	result := os.Getenv("ONEBOT_PLUGIN_PATH")
	if result != "" {
		logrus.Infof("ONEBOT_PLUGIN_PATH: %s", result)
		return result
	}
	logrus.Infof("ONEBOT_PLUGIN_PATH: %s", "./plugins")
	return "./plugins"
}

func Start(ctx context.Context) error {
	// load all plugins
	conf, err := config.LoadConfig("onebot-plus")
	if err != nil {
		return err
	}
	channelCli, err := api.NewOnebotApiClient(conf.Api)
	if err != nil {
		return err
	}
	eventCli, err := event.NewOnebotEventCli(conf.Event)
	if err != nil {
		return err
	}
	err = scanPath(
		getRootDir(),
		func(filePath string) error {
			p, rpcCli, err := LoadPlugin(filePath)
			if err != nil {
				return err
			}
			return Register(filePath, p, rpcCli)
		})
	if err != nil {
		return err
	}
	// into
	for _, p := range GetAllPlugins() {
		err := p.Plugin.Init(channelCli)
		if err != nil {
			logrus.Error(err)
		}
	}
	defer func() {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			p.Plugin.BeforeExit(channelCli)
			p.RPClient.Kill()
		}
	}()
	err = StartListenEvent(ctx, eventCli, channelCli)
	if err != nil {
		return err
	}
	return nil
}

func StartListenEvent(ctx context.Context, eventCli *event.OnebotEventClient, apiCli api.OnebotApiClientInterface) error {
	eventCli.ListenMessagePrivate(
		func(data model.EventMessagePrivate) error {
			plugins := GetAllPlugins()
			for _, p := range plugins {
				// recover
				defer func() {
					if err := recover(); err != nil {
						logrus.Error(err)
					}
				}()
				err := p.Plugin.HandleMessagePrivate(&data, apiCli)
				if err != nil {
					logrus.Error(err)
				}
			}
			return nil
		},
	)
	eventCli.ListenMessageGroup(func(data model.EventMessageGroup) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleMessageGroup(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenNoticeGroupUpload(func(data model.EventNoticeGroupUpload) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleNoticeGroupUpload(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenNoticeGroupAdmin(func(data model.EventNoticeGroupAdmin) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleNoticeGroupAdmin(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenNoticeGroupDecrease(func(data model.EventNoticeGroupDecrease) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleNoticeGroupDecrease(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenNoticeGroupIncrease(func(data model.EventNoticeGroupIncrease) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleNoticeGroupIncrease(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenNoticeGroupBan(func(data model.EventNoticeGroupBan) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleNoticeGroupBan(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenNoticeFriendAdd(func(data model.EventNoticeFriendAdd) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleNoticeFriendAdd(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenNoticeGroupRecall(func(data model.EventNoticeGroupRecall) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleNoticeGroupRecall(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenNoticeFriendRecall(func(data model.EventNoticeFriendRecall) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleNoticeFriendRecall(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenNoticeGroupNotifyPoke(func(data model.EventNoticeGroupNotifyPoke) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleNoticeGroupNotifyPoke(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenNoticeGroupNotifyHonor(func(data model.EventNoticeGroupNotifyHonor) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleNoticeGroupNotifyHonor(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenNoticeGroupNotifyLuckyKing(func(data model.EventNoticeGroupNotifyLuckyKing) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleNoticeGroupNotifyLuckyKing(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenRequestFriend(func(data model.EventRequestFriend) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleRequestFriend(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenRequestGroup(func(data model.EventRequestGroup) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleRequestGroup(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenMetaLifecycle(func(data model.EventMetaLifecycle) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleMetaLifecycle(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.ListenMetaHeartBeat(func(data model.EventMetaHeartbeat) error {
		plugins := GetAllPlugins()
		for _, p := range plugins {
			// recover
			defer func() {
				if err := recover(); err != nil {
					logrus.Error(err)
				}
			}()
			err := p.Plugin.HandleMetaHeartBeat(&data, apiCli)
			if err != nil {
				logrus.Error(err)
			}
		}
		return nil
	})
	eventCli.StartListenWithCtx(ctx)
	return nil
}

func LoadPlugin(path string) (plus_plugin.OnebotInterface, *plugin.Client, error) {
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plus_plugin.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"main": &plus_plugin.OnebotGRPCPlugin{},
		},
		Cmd: exec.Command(path),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	// defer client.Kill()
	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		logrus.Error(err)
		return nil, nil, err
	}
	// Request the plugin
	raw, err := rpcClient.Dispense("main")
	if err != nil {
		return nil, nil, err
	}
	d := raw.(plus_plugin.OnebotInterface)
	return d, client, nil
}
