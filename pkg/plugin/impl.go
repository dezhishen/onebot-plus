package plugin

import (
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
)

//插件实现
type onebotEventPluginRealImpl struct {
	//插件Id
	idCallback func() string
	//插件名称
	nameCallback func() string
	//插件描述
	descriptionCallback func() string
	//插件帮助
	helpCallback func() string
	//私聊消息
	messagePrivateCallback func(*model.EventMessagePrivate) error
	//群组消息
	messageGroupCallback func(*model.EventMessageGroup) error
	//生命周期
	metaLifecycleCallback func(*model.EventMetaLifecycle) error
	//心跳
	metaHeartbeatCallback func(*model.EventMetaHeartbeat) error
	//群文件上传
	noticeGroupUploadCallback func(*model.EventNoticeGroupUpload) error
	//群管理员变动
	noticeGroupAdminCallback func(*model.EventNoticeGroupAdmin) error
	//群成员减少
	noticeGroupDecreaseCallback func(*model.EventNoticeGroupDecrease) error
	//群成员增加
	noticeGroupIncreaseCallback func(*model.EventNoticeGroupIncrease) error
	//群禁言
	noticeGroupBanCallback func(*model.EventNoticeGroupBan) error
	//群消息撤回
	noticeGroupRecallCallback func(*model.EventNoticeGroupRecall) error
	//群内戳一戳
	noticeGroupNotifyPokeCallback func(*model.EventNoticeGroupNotifyPoke) error
	//群红包运气王
	noticeGroupNotifyLuckyKingCallback func(*model.EventNoticeGroupNotifyLuckyKing) error
	//群成员荣誉变更
	noticeGroupNotifyHonorCallback func(*model.EventNoticeGroupNotifyHonor) error
	//好友添加
	noticeFriendAddCallback func(*model.EventNoticeFriendAdd) error
	//好友消息撤回
	noticeFriendRecallCallback func(*model.EventNoticeFriendRecall) error
	//好友添加请求
	requestFriendCallback func(*model.EventRequestFriend) error
	//群添加/邀请请求
	requestGroupCallback func(*model.EventRequestGroup) error
}

func newDefaultOnebotEventPluginRealImpl() *onebotEventPluginRealImpl {
	return &onebotEventPluginRealImpl{
		//插件Id
		idCallback: returnEmptyString,
		//插件名称
		nameCallback: returnEmptyString,
		//插件描述
		descriptionCallback: returnEmptyString,
		//插件帮助
		helpCallback: returnEmptyString,
		//私聊消息
		messagePrivateCallback: func(*model.EventMessagePrivate) error { return nil },
		//群组消息
		messageGroupCallback: func(*model.EventMessageGroup) error { return nil },
		//生命周期
		metaLifecycleCallback: func(*model.EventMetaLifecycle) error { return nil },
		//心跳
		metaHeartbeatCallback: func(*model.EventMetaHeartbeat) error { return nil },
		//群文件上传
		noticeGroupUploadCallback: func(*model.EventNoticeGroupUpload) error { return nil },
		//群管理员变动
		noticeGroupAdminCallback: func(*model.EventNoticeGroupAdmin) error { return nil },
		//群成员减少
		noticeGroupDecreaseCallback: func(*model.EventNoticeGroupDecrease) error { return nil },
		//群成员增加
		noticeGroupIncreaseCallback: func(*model.EventNoticeGroupIncrease) error { return nil },
		//群禁言
		noticeGroupBanCallback: func(*model.EventNoticeGroupBan) error { return nil },
		//群消息撤回
		noticeGroupRecallCallback: func(*model.EventNoticeGroupRecall) error { return nil },
		//群内戳一戳
		noticeGroupNotifyPokeCallback: func(*model.EventNoticeGroupNotifyPoke) error { return nil },
		//群红包运气王
		noticeGroupNotifyLuckyKingCallback: func(*model.EventNoticeGroupNotifyLuckyKing) error { return nil },
		//群成员荣誉变更
		noticeGroupNotifyHonorCallback: func(*model.EventNoticeGroupNotifyHonor) error { return nil },
		//好友添加
		noticeFriendAddCallback: func(*model.EventNoticeFriendAdd) error { return nil },
		//好友消息撤回
		noticeFriendRecallCallback: func(*model.EventNoticeFriendRecall) error { return nil },
		//好友添加请求
		requestFriendCallback: func(*model.EventRequestFriend) error { return nil },
		//群添加/邀请请求
		requestGroupCallback: func(*model.EventRequestGroup) error { return nil },
	}
}

//插件Id
func (m *onebotEventPluginRealImpl) Id() string {
	return m.idCallback()
}

//插件名称
func (m *onebotEventPluginRealImpl) Name() string {
	return m.nameCallback()
}

//插件描述
func (m *onebotEventPluginRealImpl) Description() string {
	return m.descriptionCallback()
}

//插件帮助
func (m *onebotEventPluginRealImpl) Help() string {
	return m.helpCallback()
}

//私聊消息
func (m *onebotEventPluginRealImpl) MessagePrivate(req *model.EventMessagePrivate) error {
	return m.messagePrivateCallback(req)
}

//群组消息
func (m *onebotEventPluginRealImpl) MessageGroup(req *model.EventMessageGroup) error {
	return m.messageGroupCallback(req)
}

//生命周期
func (m *onebotEventPluginRealImpl) MetaLifecycle(req *model.EventMetaLifecycle) error {
	return m.metaLifecycleCallback(req)
}

//心跳
func (m *onebotEventPluginRealImpl) MetaHeartbeat(req *model.EventMetaHeartbeat) error {
	return m.metaHeartbeatCallback(req)
}

//群文件上传
func (m *onebotEventPluginRealImpl) NoticeGroupUpload(req *model.EventNoticeGroupUpload) error {
	return m.noticeGroupUploadCallback(req)
}

//群管理员变动
func (m *onebotEventPluginRealImpl) NoticeGroupAdmin(req *model.EventNoticeGroupAdmin) error {
	return m.noticeGroupAdminCallback(req)
}

//群成员减少
func (m *onebotEventPluginRealImpl) NoticeGroupDecrease(req *model.EventNoticeGroupDecrease) error {
	return m.noticeGroupDecreaseCallback(req)
}

//群成员增加
func (m *onebotEventPluginRealImpl) NoticeGroupIncrease(req *model.EventNoticeGroupIncrease) error {
	return m.noticeGroupIncreaseCallback(req)
}

//群禁言
func (m *onebotEventPluginRealImpl) NoticeGroupBan(req *model.EventNoticeGroupBan) error {
	return m.noticeGroupBanCallback(req)
}

//群消息撤回
func (m *onebotEventPluginRealImpl) NoticeGroupRecall(req *model.EventNoticeGroupRecall) error {
	return m.noticeGroupRecallCallback(req)
}

//群内戳一戳
func (m *onebotEventPluginRealImpl) NoticeGroupNotifyPoke(req *model.EventNoticeGroupNotifyPoke) error {
	return m.noticeGroupNotifyPokeCallback(req)
}

//群红包运气王
func (m *onebotEventPluginRealImpl) NoticeGroupNotifyLuckyKing(req *model.EventNoticeGroupNotifyLuckyKing) error {
	return m.noticeGroupNotifyLuckyKingCallback(req)
}

//群成员荣誉变更
func (m *onebotEventPluginRealImpl) NoticeGroupNotifyHonor(req *model.EventNoticeGroupNotifyHonor) error {
	return m.noticeGroupNotifyHonorCallback(req)
}

//好友添加
func (m *onebotEventPluginRealImpl) NoticeFriendAdd(req *model.EventNoticeFriendAdd) error {
	return m.noticeFriendAddCallback(req)
}

//好友消息撤回
func (m *onebotEventPluginRealImpl) NoticeFriendRecall(req *model.EventNoticeFriendRecall) error {
	return m.noticeFriendRecallCallback(req)
}

//好友添加请求
func (m *onebotEventPluginRealImpl) RequestFriend(req *model.EventRequestFriend) error {
	return m.requestFriendCallback(req)
}

//群添加/邀请请求
func (m *onebotEventPluginRealImpl) RequestGroup(req *model.EventRequestGroup) error {
	return m.requestGroupCallback(req)
}

//启动插件
func (p *onebotEventPluginRealImpl) Start() {
	logrus.Print("start...")
	var pluginMap = map[string]plugin.Plugin{
		"main": &onebotEventPluginGRPC{Impl: p},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
	logrus.Print("exit...")
}
