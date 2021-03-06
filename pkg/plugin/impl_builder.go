package plugin

import (
	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

//插件实现
type onebotEventPluginBuilder struct {
	impl *onebotEventPluginRealImpl
}

//插件Id
func (builder *onebotEventPluginBuilder) Id(id string) *onebotEventPluginBuilder {
	builder.impl.id = id
	return builder
}

//插件名称
func (builder *onebotEventPluginBuilder) Name(name string) *onebotEventPluginBuilder {
	builder.impl.name = name
	return builder
}

//插件描述
func (builder *onebotEventPluginBuilder) Description(description string) *onebotEventPluginBuilder {
	builder.impl.description = description
	return builder
}

//插件帮助
func (builder *onebotEventPluginBuilder) Help(help string) *onebotEventPluginBuilder {
	builder.impl.help = help
	return builder
}

//插件初始化回调
func (builder *onebotEventPluginBuilder) Init(callback func(cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.initCallback = callback
	return builder
}

//插件退出前回调
func (builder *onebotEventPluginBuilder) BeforeExit(callback func(cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.beforeExitCallback = callback
	return builder
}

//私聊消息
func (builder *onebotEventPluginBuilder) MessagePrivate(callback func(req *model.EventMessagePrivate, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.messagePrivateCallback = callback
	return builder
}

//群组消息
func (builder *onebotEventPluginBuilder) MessageGroup(callback func(req *model.EventMessageGroup, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.messageGroupCallback = callback
	return builder
}

//生命周期
func (builder *onebotEventPluginBuilder) MetaLifecycle(callback func(req *model.EventMetaLifecycle, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.metaLifecycleCallback = callback
	return builder
}

//心跳
func (builder *onebotEventPluginBuilder) MetaHeartbeat(callback func(req *model.EventMetaHeartbeat, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.metaHeartbeatCallback = callback
	return builder
}

//群文件上传
func (builder *onebotEventPluginBuilder) NoticeGroupUpload(callback func(req *model.EventNoticeGroupUpload, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.noticeGroupUploadCallback = callback
	return builder
}

//群管理员变动
func (builder *onebotEventPluginBuilder) NoticeGroupAdmin(callback func(req *model.EventNoticeGroupAdmin, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.noticeGroupAdminCallback = callback
	return builder
}

//群成员减少
func (builder *onebotEventPluginBuilder) NoticeGroupDecrease(callback func(req *model.EventNoticeGroupDecrease, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.noticeGroupDecreaseCallback = callback
	return builder
}

//群成员增加
func (builder *onebotEventPluginBuilder) NoticeGroupIncrease(callback func(req *model.EventNoticeGroupIncrease, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.noticeGroupIncreaseCallback = callback
	return builder
}

//群禁言
func (builder *onebotEventPluginBuilder) NoticeGroupBan(callback func(req *model.EventNoticeGroupBan, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.noticeGroupBanCallback = callback
	return builder
}

//群消息撤回
func (builder *onebotEventPluginBuilder) NoticeGroupRecall(callback func(req *model.EventNoticeGroupRecall, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.noticeGroupRecallCallback = callback
	return builder
}

//群内戳一戳
func (builder *onebotEventPluginBuilder) NoticeGroupNotifyPoke(callback func(req *model.EventNoticeGroupNotifyPoke, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.noticeGroupNotifyPokeCallback = callback
	return builder
}

//群红包运气王
func (builder *onebotEventPluginBuilder) NoticeGroupNotifyLuckyKing(callback func(req *model.EventNoticeGroupNotifyLuckyKing, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.noticeGroupNotifyLuckyKingCallback = callback
	return builder
}

//群成员荣誉变更
func (builder *onebotEventPluginBuilder) NoticeGroupNotifyHonor(callback func(req *model.EventNoticeGroupNotifyHonor, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.noticeGroupNotifyHonorCallback = callback
	return builder
}

//好友添加
func (builder *onebotEventPluginBuilder) NoticeFriendAdd(callback func(req *model.EventNoticeFriendAdd, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.noticeFriendAddCallback = callback
	return builder
}

//好友消息撤回
func (builder *onebotEventPluginBuilder) NoticeFriendRecall(callback func(req *model.EventNoticeFriendRecall, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.noticeFriendRecallCallback = callback
	return builder
}

//好友添加请求
func (builder *onebotEventPluginBuilder) RequestFriend(callback func(req *model.EventRequestFriend, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.requestFriendCallback = callback
	return builder
}

//群添加/邀请请求
func (builder *onebotEventPluginBuilder) RequestGroup(callback func(req *model.EventRequestGroup, cli cli.OnebotCli) error) *onebotEventPluginBuilder {
	builder.impl.requestGroupCallback = callback
	return builder
}

func (b *onebotEventPluginBuilder) Build() *onebotEventPluginRealImpl {
	return b.impl
}

//插件建造者
func NewOnebotEventPluginBuilder() *onebotEventPluginBuilder {
	result := &onebotEventPluginBuilder{
		impl: newDefaultOnebotEventPluginRealImpl(),
	}
	return result
}
