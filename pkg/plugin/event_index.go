package plugin

import (
	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotEventCallBackInterface interface {
	HandleMessagePrivate(data *model.EventMessagePrivate, onebotApi api.OnebotAPiClientInterface) error
	HandleMessageGroup(data *model.EventMessageGroup, onebotApi api.OnebotAPiClientInterface) error
	HandleNoticeGroupUpload(data *model.EventNoticeGroupUpload, onebotApi api.OnebotAPiClientInterface) error
	HandleNoticeGroupAdmin(data *model.EventNoticeGroupAdmin, onebotApi api.OnebotAPiClientInterface) error
	HandleNoticeGroupDecrease(data *model.EventNoticeGroupDecrease, onebotApi api.OnebotAPiClientInterface) error
	HandleNoticeGroupIncrease(data *model.EventNoticeGroupIncrease, onebotApi api.OnebotAPiClientInterface) error
	HandleNoticeGroupBan(data *model.EventNoticeGroupBan, onebotApi api.OnebotAPiClientInterface) error
	HandleNoticeFriendAdd(data *model.EventNoticeFriendAdd, onebotApi api.OnebotAPiClientInterface) error
	HandleNoticeGroupRecall(data *model.EventNoticeGroupRecall, onebotApi api.OnebotAPiClientInterface) error
	HandleNoticeFriendRecall(data *model.EventNoticeFriendRecall, onebotApi api.OnebotAPiClientInterface) error
	HandleNoticeGroupNotifyPoke(data *model.EventNoticeGroupNotifyPoke, onebotApi api.OnebotAPiClientInterface) error
	HandleNoticeGroupNotifyHonor(data *model.EventNoticeGroupNotifyHonor, onebotApi api.OnebotAPiClientInterface) error
	HandleNoticeGroupNotifyLuckyKing(data *model.EventNoticeGroupNotifyLuckyKing, onebotApi api.OnebotAPiClientInterface) error
	HandleRequestFriend(data *model.EventRequestFriend, onebotApi api.OnebotAPiClientInterface) error
	HandleRequestGroup(data *model.EventRequestGroup, onebotApi api.OnebotAPiClientInterface) error
	HandleMetaLifecycle(data *model.EventMetaLifecycle, onebotApi api.OnebotAPiClientInterface) error
	HandleMetaHeartBeat(data *model.EventMetaHeartbeat, onebotApi api.OnebotAPiClientInterface) error
}
