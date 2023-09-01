package plugin

import (
	"github.com/dezhishen/onebot-plus/pkg/plugin/base"
	"github.com/dezhishen/onebot-plus/pkg/plugin/event"
	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
)

type OnebotInterface interface {
	base.OnebotPluginBase
	event.OnebotEventCallBackInterface
}

type _onebotImpl struct {
	_id                               string
	_name                             string
	_description                      string
	_help                             string
	_init                             func(cli api.OnebotApiClientInterface) error
	_beforeExit                       func(cli api.OnebotApiClientInterface) error
	_handleMessagePrivate             func(data *model.EventMessagePrivate, onebotApi api.OnebotApiClientInterface) error
	_handleMessageGroup               func(data *model.EventMessageGroup, onebotApi api.OnebotApiClientInterface) error
	_handleNoticeGroupUpload          func(data *model.EventNoticeGroupUpload, onebotApi api.OnebotApiClientInterface) error
	_handleNoticeGroupAdmin           func(data *model.EventNoticeGroupAdmin, onebotApi api.OnebotApiClientInterface) error
	_handleNoticeGroupDecrease        func(data *model.EventNoticeGroupDecrease, onebotApi api.OnebotApiClientInterface) error
	_handleNoticeGroupIncrease        func(data *model.EventNoticeGroupIncrease, onebotApi api.OnebotApiClientInterface) error
	_handleNoticeGroupBan             func(data *model.EventNoticeGroupBan, onebotApi api.OnebotApiClientInterface) error
	_handleNoticeFriendAdd            func(data *model.EventNoticeFriendAdd, onebotApi api.OnebotApiClientInterface) error
	_handleNoticeGroupRecall          func(data *model.EventNoticeGroupRecall, onebotApi api.OnebotApiClientInterface) error
	_handleNoticeFriendRecall         func(data *model.EventNoticeFriendRecall, onebotApi api.OnebotApiClientInterface) error
	_handleNoticeGroupNotifyPoke      func(data *model.EventNoticeGroupNotifyPoke, onebotApi api.OnebotApiClientInterface) error
	_handleNoticeGroupNotifyHonor     func(data *model.EventNoticeGroupNotifyHonor, onebotApi api.OnebotApiClientInterface) error
	_handleNoticeGroupNotifyLuckyKing func(data *model.EventNoticeGroupNotifyLuckyKing, onebotApi api.OnebotApiClientInterface) error
	_handleRequestFriend              func(data *model.EventRequestFriend, onebotApi api.OnebotApiClientInterface) error
	_handleRequestGroup               func(data *model.EventRequestGroup, onebotApi api.OnebotApiClientInterface) error
	_handleMetaLifecycle              func(data *model.EventMetaLifecycle, onebotApi api.OnebotApiClientInterface) error
	_handleMetaHeartBeat              func(data *model.EventMetaHeartbeat, onebotApi api.OnebotApiClientInterface) error
}

// 插件Id
func (o *_onebotImpl) Id() string {
	return o._id
}

// 插件名称
func (o *_onebotImpl) Name() string {
	return o._name
}

// 插件描述
func (o *_onebotImpl) Description() string {
	return o._description
}

// 插件帮助
func (o *_onebotImpl) Help() string {
	return o._help
}

// 生命周期
func (o *_onebotImpl) Init(cli api.OnebotApiClientInterface) error {
	if o._init == nil {
		return nil
	}
	return o._init(cli)
}

// 退出前
func (o *_onebotImpl) BeforeExit(cli api.OnebotApiClientInterface) error {
	if o._beforeExit == nil {
		return nil
	}
	return o._beforeExit(cli)
}

func (o *_onebotImpl) HandleMessagePrivate(data *model.EventMessagePrivate, onebotApi api.OnebotApiClientInterface) error {
	if o._handleMessagePrivate == nil {
		return nil
	}
	return o._handleMessagePrivate(data, onebotApi)
}

func (o *_onebotImpl) HandleMessageGroup(data *model.EventMessageGroup, onebotApi api.OnebotApiClientInterface) error {
	if o._handleMessageGroup == nil {
		return nil
	}
	return o._handleMessageGroup(data, onebotApi)
}

func (o *_onebotImpl) HandleNoticeGroupUpload(data *model.EventNoticeGroupUpload, onebotApi api.OnebotApiClientInterface) error {
	if o._handleNoticeGroupUpload == nil {
		return nil
	}
	return o._handleNoticeGroupUpload(data, onebotApi)
}

func (o *_onebotImpl) HandleNoticeGroupAdmin(data *model.EventNoticeGroupAdmin, onebotApi api.OnebotApiClientInterface) error {
	if o._handleNoticeGroupAdmin == nil {
		return nil
	}
	return o._handleNoticeGroupAdmin(data, onebotApi)
}

func (o *_onebotImpl) HandleNoticeGroupDecrease(data *model.EventNoticeGroupDecrease, onebotApi api.OnebotApiClientInterface) error {
	if o._handleNoticeGroupDecrease == nil {
		return nil
	}
	return o._handleNoticeGroupDecrease(data, onebotApi)
}

func (o *_onebotImpl) HandleNoticeGroupIncrease(data *model.EventNoticeGroupIncrease, onebotApi api.OnebotApiClientInterface) error {
	if o._handleNoticeGroupIncrease == nil {
		return nil
	}
	return o._handleNoticeGroupIncrease(data, onebotApi)
}

func (o *_onebotImpl) HandleNoticeGroupBan(data *model.EventNoticeGroupBan, onebotApi api.OnebotApiClientInterface) error {
	if o._handleNoticeGroupBan == nil {
		return nil
	}
	return o._handleNoticeGroupBan(data, onebotApi)
}

func (o *_onebotImpl) HandleNoticeFriendAdd(data *model.EventNoticeFriendAdd, onebotApi api.OnebotApiClientInterface) error {
	if o._handleNoticeFriendAdd == nil {
		return nil
	}
	return o._handleNoticeFriendAdd(data, onebotApi)
}

func (o *_onebotImpl) HandleNoticeGroupRecall(data *model.EventNoticeGroupRecall, onebotApi api.OnebotApiClientInterface) error {
	if o._handleNoticeGroupRecall == nil {
		return nil
	}
	return o._handleNoticeGroupRecall(data, onebotApi)
}

func (o *_onebotImpl) HandleNoticeFriendRecall(data *model.EventNoticeFriendRecall, onebotApi api.OnebotApiClientInterface) error {
	if o._handleNoticeFriendRecall == nil {
		return nil
	}
	return o._handleNoticeFriendRecall(data, onebotApi)
}

func (o *_onebotImpl) HandleNoticeGroupNotifyPoke(data *model.EventNoticeGroupNotifyPoke, onebotApi api.OnebotApiClientInterface) error {
	if o._handleNoticeGroupNotifyPoke == nil {
		return nil
	}
	return o._handleNoticeGroupNotifyPoke(data, onebotApi)
}

func (o *_onebotImpl) HandleNoticeGroupNotifyHonor(data *model.EventNoticeGroupNotifyHonor, onebotApi api.OnebotApiClientInterface) error {
	if o._handleNoticeGroupNotifyHonor == nil {
		return nil
	}
	return o._handleNoticeGroupNotifyHonor(data, onebotApi)
}

func (o *_onebotImpl) HandleNoticeGroupNotifyLuckyKing(data *model.EventNoticeGroupNotifyLuckyKing, onebotApi api.OnebotApiClientInterface) error {
	if o._handleNoticeGroupNotifyLuckyKing == nil {
		return nil
	}
	return o._handleNoticeGroupNotifyLuckyKing(data, onebotApi)
}

func (o *_onebotImpl) HandleRequestFriend(data *model.EventRequestFriend, onebotApi api.OnebotApiClientInterface) error {
	if o._handleRequestFriend == nil {
		return nil
	}
	return o._handleRequestFriend(data, onebotApi)
}

func (o *_onebotImpl) HandleRequestGroup(data *model.EventRequestGroup, onebotApi api.OnebotApiClientInterface) error {
	if o._handleRequestGroup == nil {
		return nil
	}
	return o._handleRequestGroup(data, onebotApi)
}

func (o *_onebotImpl) HandleMetaLifecycle(data *model.EventMetaLifecycle, onebotApi api.OnebotApiClientInterface) error {
	if o._handleMetaLifecycle == nil {
		return nil
	}
	return o._handleMetaLifecycle(data, onebotApi)
}

func (o *_onebotImpl) HandleMetaHeartBeat(data *model.EventMetaHeartbeat, onebotApi api.OnebotApiClientInterface) error {
	if o._handleMetaHeartBeat == nil {
		return nil
	}
	return o._handleMetaHeartBeat(data, onebotApi)
}

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "onebot-plus",
	MagicCookieValue: "onebot-plus",
}

func (o *_onebotImpl) Start() {
	logrus.Info("start...")
	var pluginMap = map[string]plugin.Plugin{
		"main": &OnebotGRPCPlugin{Impl: o},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
	logrus.Info("exit...")
}
