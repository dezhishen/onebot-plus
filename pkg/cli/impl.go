package cli

import (
	sdk_api "github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotCliRealImpl struct {
}

//发送消息
func (d *OnebotCliRealImpl) SendMsg(msg *model.MsgForSend) (int64, error) {
	return sdk_api.SendMsg(msg)
}

//发送私聊消息
func (d *OnebotCliRealImpl) SendPrivateMsg(msg *model.PrivateMsg) (int64, error) {
	return sdk_api.SendPrivateMsg(msg)

}

// 发送群消息
func (d *OnebotCliRealImpl) SendGroupMsg(msg *model.GroupMsg) (int64, error) {
	return sdk_api.SendGroupMsg(msg)

}

//删除消息
func (d *OnebotCliRealImpl) DelMsg(id int64) error {
	return sdk_api.DelMsg(id)

}

//获取消息
func (d *OnebotCliRealImpl) GetMsg(id int64) (*model.MessageData, error) {
	return sdk_api.GetMsg(id)

}

//获取转发的消息
func (d *OnebotCliRealImpl) GetForwardMsg(id int64) (*model.ForwardMessageData, error) {
	return sdk_api.GetForwardMsg(id)
}

//获取登录信息
func (d *OnebotCliRealImpl) GetLoginInfo() (*model.Account, error) {
	return sdk_api.GetLoginInfo()
}

//获取陌生人信息
func (d *OnebotCliRealImpl) GetStrangerInfo(userId int64, noCache bool) (*model.Account, error) {
	return sdk_api.GetStrangerInfo(userId, noCache)
}

func (d *OnebotCliRealImpl) GetCookies(domin string) (*model.Cokies, error) {
	return sdk_api.GetCookies(domin)
}

func (d *OnebotCliRealImpl) GetCSRFToken() (*model.CSRFToken, error) {
	return sdk_api.GetCSRFToken()
}

func (d *OnebotCliRealImpl) GetCredentials(domin string) (*model.Credentials, error) {
	return sdk_api.GetCredentials(domin)
}

//获取语音
func (d *OnebotCliRealImpl) GetRecord(file string, out_format string) (*model.File, error) {
	return sdk_api.GetRecord(file, out_format)
}

//获取图片
func (d *OnebotCliRealImpl) GetImage(file string) (*model.File, error) {
	return sdk_api.GetImage(file)
}
