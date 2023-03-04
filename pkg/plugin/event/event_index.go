package event

import (
	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotEventCallBackInterface interface {
	HandleMessagePrivate(data *model.EventMessagePrivate, onebotApi api.OnebotApiClientInterface) error
	HandleMessageGroup(data *model.EventMessageGroup, onebotApi api.OnebotApiClientInterface) error
	HandleNoticeGroupUpload(data *model.EventNoticeGroupUpload, onebotApi api.OnebotApiClientInterface) error
	HandleNoticeGroupAdmin(data *model.EventNoticeGroupAdmin, onebotApi api.OnebotApiClientInterface) error
	HandleNoticeGroupDecrease(data *model.EventNoticeGroupDecrease, onebotApi api.OnebotApiClientInterface) error
	HandleNoticeGroupIncrease(data *model.EventNoticeGroupIncrease, onebotApi api.OnebotApiClientInterface) error
	HandleNoticeGroupBan(data *model.EventNoticeGroupBan, onebotApi api.OnebotApiClientInterface) error
	HandleNoticeFriendAdd(data *model.EventNoticeFriendAdd, onebotApi api.OnebotApiClientInterface) error
	HandleNoticeGroupRecall(data *model.EventNoticeGroupRecall, onebotApi api.OnebotApiClientInterface) error
	HandleNoticeFriendRecall(data *model.EventNoticeFriendRecall, onebotApi api.OnebotApiClientInterface) error
	HandleNoticeGroupNotifyPoke(data *model.EventNoticeGroupNotifyPoke, onebotApi api.OnebotApiClientInterface) error
	HandleNoticeGroupNotifyHonor(data *model.EventNoticeGroupNotifyHonor, onebotApi api.OnebotApiClientInterface) error
	HandleNoticeGroupNotifyLuckyKing(data *model.EventNoticeGroupNotifyLuckyKing, onebotApi api.OnebotApiClientInterface) error
	HandleRequestFriend(data *model.EventRequestFriend, onebotApi api.OnebotApiClientInterface) error
	HandleRequestGroup(data *model.EventRequestGroup, onebotApi api.OnebotApiClientInterface) error
	HandleMetaLifecycle(data *model.EventMetaLifecycle, onebotApi api.OnebotApiClientInterface) error
	HandleMetaHeartBeat(data *model.EventMetaHeartbeat, onebotApi api.OnebotApiClientInterface) error
}
