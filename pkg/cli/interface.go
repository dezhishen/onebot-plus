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
}
