package api

import (
	onebotPlugin "github.com/dezhishen/onebot-plus/pkg/plugin"
	sdk_api "github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
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

func StartMessageAPI() {
	cli := &MessageCliRealImpl{}
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"main": &MessageCliGRPCPlugin{Impl: cli},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: onebotPlugin.HandshakeConfig,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}
