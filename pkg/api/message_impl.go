package api

import (
	sdk_api "github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type MessageCliRealImpl struct {
}

//发送消息
func (d MessageCliRealImpl) SendMsg(msg *model.MsgForSend) (int64, error) {
	return sdk_api.SendMsg(msg)
}

//发送私聊消息
func (d MessageCliRealImpl) SendPrivateMsg(msg *model.PrivateMsg) (int64, error) {
	return sdk_api.SendPrivateMsg(msg)

}

// 发送群消息
func (d MessageCliRealImpl) SendGroupMsg(msg *model.GroupMsg) (int64, error) {
	return sdk_api.SendGroupMsg(msg)

}

//删除消息
func (d MessageCliRealImpl) DelMsg(id int64) error {
	return sdk_api.DelMsg(id)

}

//获取消息
func (d MessageCliRealImpl) GetMsg(id int64) (*model.MessageData, error) {
	return sdk_api.GetMsg(id)

}

//获取转发的消息
func (d MessageCliRealImpl) GetForwardMsg(id int64) (*model.ForwardMessageData, error) {
	return sdk_api.GetForwardMsg(id)
}
