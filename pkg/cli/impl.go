package cli

import (
	sdk_api "github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotCliRealImpl struct {
}

//发送消息
func (d *OnebotCliRealImpl) SendMsg(msg *model.MsgForSend) (*model.SendMessageResult, error) {
	return sdk_api.SendMsg(msg)
}

//发送私聊消息
func (d *OnebotCliRealImpl) SendPrivateMsg(msg *model.PrivateMsg) (*model.SendMessageResult, error) {
	return sdk_api.SendPrivateMsg(msg)

}

// 发送群消息
func (d *OnebotCliRealImpl) SendGroupMsg(msg *model.GroupMsg) (*model.SendMessageResult, error) {
	return sdk_api.SendGroupMsg(msg)

}

//转发合并消息(群)
func (d *OnebotCliRealImpl) SendGroupForwardMessage(groupId int64, messages []*model.MessageSegment) (*model.SendGroupForwardMessageDataResult, error) {
	return sdk_api.SendGroupForwardMsg(groupId, messages)
}

//转发合并消息到群
func (d *OnebotCliRealImpl) SendGroupForwardMessageByRawMsg(groupId, userId int64, name string, messages []*model.MessageSegment) (*model.SendGroupForwardMessageDataResult, error) {
	return sdk_api.SendGroupForwardMsgByRawMsg(groupId, userId, name, messages)
}

//删除消息
func (d *OnebotCliRealImpl) DelMsg(id int64) error {
	return sdk_api.DelMsg(id)

}

//获取消息
func (d *OnebotCliRealImpl) GetMsg(id int64) (*model.MessageDataResult, error) {
	return sdk_api.GetMsg(id)

}

//获取转发的消息
func (d *OnebotCliRealImpl) GetForwardMsg(id int64) (*model.ForwardMessageDataResult, error) {
	return sdk_api.GetForwardMsg(id)
}

//获取登录信息
func (d *OnebotCliRealImpl) GetLoginInfo() (*model.AccountResult, error) {
	return sdk_api.GetLoginInfo()
}

//获取陌生人信息
func (d *OnebotCliRealImpl) GetStrangerInfo(userId int64, noCache bool) (*model.AccountResult, error) {
	return sdk_api.GetStrangerInfo(userId, noCache)
}

func (d *OnebotCliRealImpl) GetCookies(domin string) (*model.CokiesResult, error) {
	return sdk_api.GetCookies(domin)
}

func (d *OnebotCliRealImpl) GetCSRFToken() (*model.CSRFTokenResult, error) {
	return sdk_api.GetCSRFToken()
}

func (d *OnebotCliRealImpl) GetCredentials(domin string) (*model.CredentialsResult, error) {
	return sdk_api.GetCredentials(domin)
}

//获取语音
func (d *OnebotCliRealImpl) GetRecord(file string, out_format string) (*model.FileResult, error) {
	return sdk_api.GetRecord(file, out_format)
}

//获取图片
func (d *OnebotCliRealImpl) GetImage(file string) (*model.FileResult, error) {
	return sdk_api.GetImage(file)
}

func (d *OnebotCliRealImpl) SendLike(userId int64, times int64) error {
	return sdk_api.SendLike(userId, times)
}

//处理加好友请求
func (d *OnebotCliRealImpl) SetFriendAddRequest(flag string, approve bool, remark string) error {
	return sdk_api.SetFriendAddRequest(flag, approve, remark)
}
func (d *OnebotCliRealImpl) GetFriendList() (*model.FriendListResult, error) {
	return sdk_api.GetFriendList()
}

// 群组踢人
func (d *OnebotCliRealImpl) SetGroupKick(groupId, userId int64, rejectAddRequest bool) error {
	return sdk_api.SetGroupKick(groupId, userId, rejectAddRequest)
}

// 群组禁言
func (d *OnebotCliRealImpl) SetGroupBan(groupId, userId, duration int64) error {
	return sdk_api.SetGroupBan(groupId, userId, duration)
}

// 群组匿名用户禁言
func (d *OnebotCliRealImpl) SetGroupAnonymousBan(groupId, duration int64, anonymousFlag string) error {
	return sdk_api.SetGroupAnonymousBan(groupId, duration, anonymousFlag)
}

//群组全员禁言
func (d *OnebotCliRealImpl) SetGroupWholeBan(groupId int64, enable bool) error {
	return sdk_api.SetGroupWholeBan(groupId, enable)
}

//群组设置管理员
func (d *OnebotCliRealImpl) SetGroupAdmin(groupId, userId int64, enable bool) error {
	return sdk_api.SetGroupAdmin(groupId, userId, enable)
}

//群组匿名
func (d *OnebotCliRealImpl) SetGroupAnonymous(groupId int64, enable bool) error {
	return sdk_api.SetGroupAnonymous(groupId, enable)
}

//设置群名片（群备注）
func (d *OnebotCliRealImpl) SetGroupCard(groupId, userId int64, card string) error {
	return sdk_api.SetGroupCard(groupId, userId, card)
}

//设置群名
func (d *OnebotCliRealImpl) SetGroupName(groupId int64, groupName string) error {
	return sdk_api.SetGroupName(groupId, groupName)
}

//退出群组
func (d *OnebotCliRealImpl) SetGroupLeave(groupId int64, isDismiss bool) error {
	return sdk_api.SetGroupLeave(groupId, isDismiss)
}

//设置群组专属头衔
func (d *OnebotCliRealImpl) SetGroupSpecialTitle(groupId, userId, duration int64, specialTitle string) error {
	return sdk_api.SetGroupSpecialTitle(groupId, userId, duration, specialTitle)
}

//处理加群请求／邀请
func (d *OnebotCliRealImpl) SetGroupAddRequest(flag, subType, reason string, approve bool) error {
	return sdk_api.SetGroupAddRequest(flag, subType, reason, approve)
}

//获取群信息
func (d *OnebotCliRealImpl) GetGroupInfo(groupId int64, noCache bool) (*model.GroupResult, error) {
	return sdk_api.GetGroupInfo(groupId, noCache)
}

//获取群列表
func (d *OnebotCliRealImpl) GetGroupList() (*model.GroupListResult, error) {
	return sdk_api.GetGroupList()
}

//获取群成员信息
func (d *OnebotCliRealImpl) GetGroupMemberInfo(groupId, userId int64, noCache bool) (*model.GroupMemberResult, error) {
	return sdk_api.GetGroupMemberInfo(groupId, userId, noCache)
}

//获取群成员列表
func (d *OnebotCliRealImpl) GetGroupMemberListInfo() (*model.GroupMemberListResult, error) {
	return sdk_api.GetGroupMemberListInfo()
}

//获取群荣誉信息
func (d *OnebotCliRealImpl) GetGroupHonorInfo(groupId int64, honorType string) (*model.GroupHonorInfoResult, error) {
	return sdk_api.GetGroupHonorInfo(groupId, honorType)
}

func (d *OnebotCliRealImpl) CleanCache() error {
	return sdk_api.CleanCache()
}

func (d *OnebotCliRealImpl) SetRestart(delay int64) error {
	return sdk_api.SetRestart(delay)
}

func (d *OnebotCliRealImpl) GetVersionInfo() (*model.VersionInfoResult, error) {
	return sdk_api.GetVersionInfo()
}

func (d *OnebotCliRealImpl) GetStatus() (*model.StatusInfoResult, error) {
	return sdk_api.GetStatus()
}

func (d *OnebotCliRealImpl) CanSendImage() (*model.BoolYesOfResult, error) {
	return sdk_api.CanSendImage()
}

func (d *OnebotCliRealImpl) CanSendRecord() (*model.BoolYesOfResult, error) {
	return sdk_api.CanSendRecord()
}
