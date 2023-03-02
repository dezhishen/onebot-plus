package tools

// import (
// 	"github.com/dezhishen/onebot-sdk/pkg/api"
// 	"github.com/dezhishen/onebot-sdk/pkg/model"
// ) // func Forward(cli *onebot.)

// func CreateForward(msg []*model.MessageSegment) (int64, error) {
// 	accountResult, err := api.GetLoginInfo()
// 	if err != nil {
// 		return 0, err
// 	}
// 	sendResult, err := api.SendPrivateMsg(&model.PrivateMsg{
// 		UserId:  accountResult.Data.UserId,
// 		Message: msg,
// 	})
// 	if err != nil {
// 		return 0, err
// 	}
// 	return sendResult.Data.MessageId, nil
// }
