package plugin

import (
	sdk_api "github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type onebotPluginBuilder struct {
	*_onebotImpl
}

func OnebotPluginBuilder() *onebotPluginBuilder {
	return &onebotPluginBuilder{
		_onebotImpl: &_onebotImpl{},
	}
}

func (b *onebotPluginBuilder) Id(id string) *onebotPluginBuilder {
	b._id = id
	return b
}

func (b *onebotPluginBuilder) Name(name string) *onebotPluginBuilder {
	b._name = name
	return b
}

func (b *onebotPluginBuilder) Description(description string) *onebotPluginBuilder {
	b._description = description
	return b
}

func (b *onebotPluginBuilder) Help(help string) *onebotPluginBuilder {
	b._help = help
	return b
}

func (b *onebotPluginBuilder) Init(init func(onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._init = init
	return b
}

func (b *onebotPluginBuilder) BeforeExit(beforeExit func(onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._beforeExit = beforeExit
	return b
}

// _handleMessagePrivate             func(data *model.EventMessagePrivate, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleMessagePrivate(handleMessagePrivate func(data *model.EventMessagePrivate, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleMessagePrivate = handleMessagePrivate
	return b
}

// _handleMessageGroup               func(data *model.EventMessageGroup, onebotApi api.OnebotApiClientInterface)

func (b *onebotPluginBuilder) HandleMessageGroup(handleMessageGroup func(data *model.EventMessageGroup, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleMessageGroup = handleMessageGroup
	return b
}

// _handleNoticeGroupUpload          func(data *model.EventNoticeGroupUpload, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleNoticeGroupUpload(handleNoticeGroupUpload func(data *model.EventNoticeGroupUpload, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleNoticeGroupUpload = handleNoticeGroupUpload
	return b
}

// _handleNoticeGroupAdmin           func(data *model.EventNoticeGroupAdmin, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleNoticeGroupAdmin(handleNoticeGroupAdmin func(data *model.EventNoticeGroupAdmin, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleNoticeGroupAdmin = handleNoticeGroupAdmin
	return b
}

// _handleNoticeGroupDecrease        func(data *model.EventNoticeGroupDecrease, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleNoticeGroupDecrease(handleNoticeGroupDecrease func(data *model.EventNoticeGroupDecrease, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleNoticeGroupDecrease = handleNoticeGroupDecrease
	return b
}

// _handleNoticeGroupIncrease        func(data *model.EventNoticeGroupIncrease, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleNoticeGroupIncrease(handleNoticeGroupIncrease func(data *model.EventNoticeGroupIncrease, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleNoticeGroupIncrease = handleNoticeGroupIncrease
	return b
}

// _handleNoticeGroupBan             func(data *model.EventNoticeGroupBan, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleNoticeGroupBan(handleNoticeGroupBan func(data *model.EventNoticeGroupBan, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleNoticeGroupBan = handleNoticeGroupBan
	return b
}

// _handleNoticeFriendAdd            func(data *model.EventNoticeFriendAdd, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleNoticeFriendAdd(handleNoticeFriendAdd func(data *model.EventNoticeFriendAdd, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleNoticeFriendAdd = handleNoticeFriendAdd
	return b
}

// _handleNoticeGroupRecall          func(data *model.EventNoticeGroupRecall, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleNoticeGroupRecall(handleNoticeGroupRecall func(data *model.EventNoticeGroupRecall, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleNoticeGroupRecall = handleNoticeGroupRecall
	return b
}

// _handleNoticeFriendRecall         func(data *model.EventNoticeFriendRecall, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleNoticeFriendRecall(handleNoticeFriendRecall func(data *model.EventNoticeFriendRecall, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleNoticeFriendRecall = handleNoticeFriendRecall
	return b
}

// _handleNoticeGroupNotifyPoke      func(data *model.EventNoticeGroupNotifyPoke, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleNoticeGroupNotifyPoke(handleNoticeGroupNotifyPoke func(data *model.EventNoticeGroupNotifyPoke, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleNoticeGroupNotifyPoke = handleNoticeGroupNotifyPoke
	return b
}

// _handleNoticeGroupNotifyHonor     func(data *model.EventNoticeGroupNotifyHonor, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleNoticeGroupNotifyHonor(handleNoticeGroupNotifyHonor func(data *model.EventNoticeGroupNotifyHonor, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleNoticeGroupNotifyHonor = handleNoticeGroupNotifyHonor
	return b
}

// _handleNoticeGroupNotifyLuckyKing func(data *model.EventNoticeGroupNotifyLuckyKing, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleNoticeGroupNotifyLuckyKing(handleNoticeGroupNotifyLuckyKing func(data *model.EventNoticeGroupNotifyLuckyKing, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleNoticeGroupNotifyLuckyKing = handleNoticeGroupNotifyLuckyKing
	return b
}

// _handleRequestFriend              func(data *model.EventRequestFriend, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleRequestFriend(handleRequestFriend func(data *model.EventRequestFriend, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleRequestFriend = handleRequestFriend
	return b
}

// _handleRequestGroup               func(data *model.EventRequestGroup, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleRequestGroup(handleRequestGroup func(data *model.EventRequestGroup, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleRequestGroup = handleRequestGroup
	return b
}

// _handleMetaLifecycle              func(data *model.EventMetaLifecycle, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleMetaLifecycle(handleMetaLifecycle func(data *model.EventMetaLifecycle, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleMetaLifecycle = handleMetaLifecycle
	return b
}

// _handleMetaHeartBeat              func(data *model.EventMetaHeartbeat, onebotApi api.OnebotApiClientInterface) error
func (b *onebotPluginBuilder) HandleMetaHeartBeat(handleMetaHeartBeat func(data *model.EventMetaHeartbeat, onebotApi sdk_api.OnebotApiClientInterface) error) *onebotPluginBuilder {
	b._handleMetaHeartBeat = handleMetaHeartBeat
	return b
}

func (b *onebotPluginBuilder) Build() *_onebotImpl {
	return b._onebotImpl
}
