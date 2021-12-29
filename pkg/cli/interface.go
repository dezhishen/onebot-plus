package cli

import "github.com/dezhishen/onebot-sdk/pkg/model"

type OnebotCli interface {
	//发送消息
	SendMsg(msg *model.MsgForSend) (int64, error)
	//发送私聊消息
	SendPrivateMsg(msg *model.PrivateMsg) (int64, error)
	// 发送群消息
	SendGroupMsg(msg *model.GroupMsg) (int64, error)
	//删除消息
	DelMsg(id int64) error
	//获取消息
	GetMsg(id int64) (*model.MessageData, error)
	//获取转发的消息
	GetForwardMsg(id int64) (*model.ForwardMessageData, error)
	//获取登录信息
	GetLoginInfo() (*model.Account, error)
	//获取陌生人信息
	GetStrangerInfo(userId int64, noCache bool) (*model.Account, error)

	GetCookies(domin string) (*model.Cokies, error)

	GetCSRFToken() (*model.CSRFToken, error)

	GetCredentials(domin string) (*model.Credentials, error)
	//获取语音
	GetRecord(file string, out_format string) (*model.File, error)
	//获取图片
	GetImage(file string) (*model.File, error)

	SendLike(userId int64, times int64) error
	//处理加好友请求
	SetFriendAddRequest(flag string, approve bool, remark string) error
	GetFriendList() ([]*model.Account, error)

	// 群组踢人
	SetGroupKick(groupId, userId int64, rejectAddRequest bool) error

	// 群组禁言
	SetGroupBan(groupId, userId, duration int64) error

	// 群组匿名用户禁言
	SetGroupAnonymousBan(groupId, duration int64, anonymousFlag string) error

	//群组全员禁言
	SetGroupWholeBan(groupId int64, enable bool) error

	//群组设置管理员
	SetGroupAdmin(groupId, userId int64, enable bool) error
	//群组匿名
	SetGroupAnonymous(groupId int64, enable bool) error
	//设置群名片（群备注）
	SetGroupCard(groupId, userId int64, card string) error

	//设置群名
	SetGroupName(groupId int64, groupName string) error

	//退出群组
	SetGroupLeave(groupId int64, isDismiss bool) error
	//设置群组专属头衔
	SetGroupSpecialTitle(groupId, userId, duration int64, specialTitle string) error

	//处理加群请求／邀请
	SetGroupAddRequest(flag, subType, reason string, approve bool) error

	//获取群信息
	GetGroupInfo(groupId int64, noCache bool) (*model.Group, error)

	//获取群列表
	GetGroupList() ([]*model.Group, error)

	//获取群成员信息
	GetGroupMemberInfo(groupId, userId int64, noCache bool) (*model.GroupMember, error)

	//获取群成员列表
	GetGroupMemberListInfo() ([]*model.GroupMember, error)

	//获取群荣誉信息
	GetGroupHonorInfo(groupId int64, honorType string) (*model.GroupHonorInfo, error)

	CleanCache() error

	SetRestart(delay int64) error

	GetVersionInfo() (*model.VersionInfoData, error)

	GetStatus() (*model.StatusInfoData, error)

	CanSendImage() (bool, error)

	CanSendRecord() (bool, error)
}
