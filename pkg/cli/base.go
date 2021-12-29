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
}

type OneBotCliRealImpl struct {
	MessageCli MessageCli
}
