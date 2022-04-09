package plugin

import (
	"github.com/dezhishen/onebot-plus/pkg/cli"
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
	initCallback func(cli.OnebotCli) error
	//是否禁用某个事件
	isEventDisabledSetting map[PluginEventType]bool
	//私聊消息
	messagePrivateCallback func(*model.EventMessagePrivate, cli.OnebotCli) error
	//群组消息
	messageGroupCallback func(*model.EventMessageGroup, cli.OnebotCli) error
	//生命周期
	metaLifecycleCallback func(*model.EventMetaLifecycle, cli.OnebotCli) error
	//心跳
	metaHeartbeatCallback func(*model.EventMetaHeartbeat, cli.OnebotCli) error
	//群文件上传
	noticeGroupUploadCallback func(*model.EventNoticeGroupUpload, cli.OnebotCli) error
	//群管理员变动
	noticeGroupAdminCallback func(*model.EventNoticeGroupAdmin, cli.OnebotCli) error
	//群成员减少
	noticeGroupDecreaseCallback func(*model.EventNoticeGroupDecrease, cli.OnebotCli) error
	//群成员增加
	noticeGroupIncreaseCallback func(*model.EventNoticeGroupIncrease, cli.OnebotCli) error
	//群禁言
	noticeGroupBanCallback func(*model.EventNoticeGroupBan, cli.OnebotCli) error
	//群消息撤回
	noticeGroupRecallCallback func(*model.EventNoticeGroupRecall, cli.OnebotCli) error
	//群内戳一戳
	noticeGroupNotifyPokeCallback func(*model.EventNoticeGroupNotifyPoke, cli.OnebotCli) error
	//群红包运气王
	noticeGroupNotifyLuckyKingCallback func(*model.EventNoticeGroupNotifyLuckyKing, cli.OnebotCli) error
	//群成员荣誉变更
	noticeGroupNotifyHonorCallback func(*model.EventNoticeGroupNotifyHonor, cli.OnebotCli) error
	//好友添加
	noticeFriendAddCallback func(*model.EventNoticeFriendAdd, cli.OnebotCli) error
	//好友消息撤回
	noticeFriendRecallCallback func(*model.EventNoticeFriendRecall, cli.OnebotCli) error
	//好友添加请求
	requestFriendCallback func(*model.EventRequestFriend, cli.OnebotCli) error
	//群添加/邀请请求
	requestGroupCallback func(*model.EventRequestGroup, cli.OnebotCli) error
	//退出前回调
	beforeExitCallback func(cli.OnebotCli) error
}

func newDefaultOnebotEventPluginRealImpl() *onebotEventPluginRealImpl {
	isEventDisabled := make(map[PluginEventType]bool)
	return &onebotEventPluginRealImpl{
		//是否禁用某个事件
		isEventDisabledSetting: isEventDisabled,
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

//插件初始化
func (m *onebotEventPluginRealImpl) Init(onebotCli cli.OnebotCli) error {
	if m.initCallback == nil {
		return nil
	}
	return m.initCallback(onebotCli)
}

//退出前事件
func (m *onebotEventPluginRealImpl) BeforeExit(cli cli.OnebotCli) error {
	if m.beforeExitCallback == nil {
		return nil
	}
	return m.beforeExitCallback(cli)
}

//私聊消息
func (m *onebotEventPluginRealImpl) MessagePrivate(req *model.EventMessagePrivate, cli cli.OnebotCli) error {
	if m.messagePrivateCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventMessagePrivate) {
		return nil
	}
	return m.messagePrivateCallback(req, cli)
}

//群组消息
func (m *onebotEventPluginRealImpl) MessageGroup(req *model.EventMessageGroup, cli cli.OnebotCli) error {
	if m.messageGroupCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventMessageGroup) {
		return nil
	}
	return m.messageGroupCallback(req, cli)
}

//生命周期
func (m *onebotEventPluginRealImpl) MetaLifecycle(req *model.EventMetaLifecycle, cli cli.OnebotCli) error {
	if m.metaLifecycleCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventMetaLifecycle) {
		return nil
	}
	return m.metaLifecycleCallback(req, cli)
}

//心跳
func (m *onebotEventPluginRealImpl) MetaHeartbeat(req *model.EventMetaHeartbeat, cli cli.OnebotCli) error {
	if m.metaHeartbeatCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventMetaHeartbeat) {
		return nil
	}
	return m.metaHeartbeatCallback(req, cli)
}

//群文件上传
func (m *onebotEventPluginRealImpl) NoticeGroupUpload(req *model.EventNoticeGroupUpload, cli cli.OnebotCli) error {
	if m.noticeGroupUploadCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventNoticeGroupUpload) {
		return nil
	}
	return m.noticeGroupUploadCallback(req, cli)
}

//群管理员变动
func (m *onebotEventPluginRealImpl) NoticeGroupAdmin(req *model.EventNoticeGroupAdmin, cli cli.OnebotCli) error {
	if m.noticeGroupAdminCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventNoticeGroupAdmin) {
		return nil
	}
	return m.noticeGroupAdminCallback(req, cli)
}

//群成员减少
func (m *onebotEventPluginRealImpl) NoticeGroupDecrease(req *model.EventNoticeGroupDecrease, cli cli.OnebotCli) error {
	if m.noticeGroupDecreaseCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventNoticeGroupDecrease) {
		return nil
	}
	return m.noticeGroupDecreaseCallback(req, cli)
}

//群成员增加
func (m *onebotEventPluginRealImpl) NoticeGroupIncrease(req *model.EventNoticeGroupIncrease, cli cli.OnebotCli) error {
	if m.noticeGroupIncreaseCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventNoticeGroupIncrease) {
		return nil
	}
	return m.noticeGroupIncreaseCallback(req, cli)
}

//群禁言
func (m *onebotEventPluginRealImpl) NoticeGroupBan(req *model.EventNoticeGroupBan, cli cli.OnebotCli) error {
	if m.noticeGroupBanCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventNoticeGroupBan) {
		return nil
	}
	return m.noticeGroupBanCallback(req, cli)
}

//群消息撤回
func (m *onebotEventPluginRealImpl) NoticeGroupRecall(req *model.EventNoticeGroupRecall, cli cli.OnebotCli) error {
	if m.noticeGroupRecallCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventNoticeGroupRecall) {
		return nil
	}
	return m.noticeGroupRecallCallback(req, cli)
}

//群内戳一戳
func (m *onebotEventPluginRealImpl) NoticeGroupNotifyPoke(req *model.EventNoticeGroupNotifyPoke, cli cli.OnebotCli) error {
	if m.noticeGroupNotifyPokeCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventNoticeGroupNotifyPoke) {
		return nil
	}
	return m.noticeGroupNotifyPokeCallback(req, cli)
}

//群红包运气王
func (m *onebotEventPluginRealImpl) NoticeGroupNotifyLuckyKing(req *model.EventNoticeGroupNotifyLuckyKing, cli cli.OnebotCli) error {
	if m.noticeGroupNotifyLuckyKingCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventNoticeGroupNotifyLuckyKing) {
		return nil
	}
	return m.noticeGroupNotifyLuckyKingCallback(req, cli)
}

//群成员荣誉变更
func (m *onebotEventPluginRealImpl) NoticeGroupNotifyHonor(req *model.EventNoticeGroupNotifyHonor, cli cli.OnebotCli) error {
	if m.noticeGroupNotifyHonorCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventNoticeGroupNotifyHonor) {
		return nil
	}
	return m.noticeGroupNotifyHonorCallback(req, cli)
}

//好友添加
func (m *onebotEventPluginRealImpl) NoticeFriendAdd(req *model.EventNoticeFriendAdd, cli cli.OnebotCli) error {
	if m.noticeFriendAddCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventNoticeFriendAdd) {
		return nil
	}
	return m.noticeFriendAddCallback(req, cli)
}

//好友消息撤回
func (m *onebotEventPluginRealImpl) NoticeFriendRecall(req *model.EventNoticeFriendRecall, cli cli.OnebotCli) error {
	if m.noticeFriendRecallCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventNoticeFriendRecall) {
		return nil
	}
	return m.noticeFriendRecallCallback(req, cli)
}

//好友添加请求
func (m *onebotEventPluginRealImpl) RequestFriend(req *model.EventRequestFriend, cli cli.OnebotCli) error {
	if m.requestFriendCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventRequestFriend) {
		return nil
	}
	return m.requestFriendCallback(req, cli)
}

//群添加/邀请请求
func (m *onebotEventPluginRealImpl) RequestGroup(req *model.EventRequestGroup, cli cli.OnebotCli) error {
	if m.requestGroupCallback == nil {
		return nil
	}
	if m.isEventDisabled(PluginEventRequestGroup) {
		return nil
	}
	return m.requestGroupCallback(req, cli)
}

//启动插件
func (p *onebotEventPluginRealImpl) Start() {
	logrus.Info("start...")
	var pluginMap = map[string]plugin.Plugin{
		"main": &onebotEventPluginGRPC{Impl: p},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
	logrus.Info("exit...")
}

func (m *onebotEventPluginRealImpl) isEventDisabled(key PluginEventType) bool {
	v, ok := m.isEventDisabledSetting[PluginEventMetaLifecycle]
	if ok && v {
		return true
	}
	return false
}
