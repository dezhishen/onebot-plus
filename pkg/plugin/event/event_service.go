package event

import (
	"context"

	"github.com/dezhishen/onebot-plus/pkg/plugin/api"
	"github.com/hashicorp/go-plugin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OnebotEventCallbackServerStub struct {
	Broker *plugin.GRPCBroker
	// 具体实现，仅当业务接口实现基于Go时该字段有用
	Impl OnebotEventCallBackInterface
}

// HandleMessagePrivate(data *model.EventMessagePrivate, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleMessagePrivate(ctx context.Context, in *EventMessagePrivateGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleMessagePrivate(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleMessageGroup(data *model.EventMessageGroup, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleMessageGroup(ctx context.Context, in *EventMessageGroupGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleMessageGroup(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleNoticeGroupUpload(data *model.EventNoticeGroupUpload, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleNoticeGroupUpload(ctx context.Context, in *EventNoticeGroupUploadGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleNoticeGroupUpload(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleNoticeGroupAdmin(data *model.EventNoticeGroupAdmin, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleNoticeGroupAdmin(ctx context.Context, in *EventNoticeGroupAdminGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleNoticeGroupAdmin(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleNoticeGroupDecrease(data *model.EventNoticeGroupDecrease, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleNoticeGroupDecrease(ctx context.Context, in *EventNoticeGroupDecreaseGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleNoticeGroupDecrease(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleNoticeGroupIncrease(data *model.EventNoticeGroupIncrease, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleNoticeGroupIncrease(ctx context.Context, in *EventNoticeGroupIncreaseGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleNoticeGroupIncrease(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e

}

// HandleNoticeGroupBan(data *model.EventNoticeGroupBan, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleNoticeGroupBan(ctx context.Context, in *EventNoticeGroupBanGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleNoticeGroupBan(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleNoticeFriendAdd(data *model.EventNoticeFriendAdd, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleNoticeFriendAdd(ctx context.Context, in *EventNoticeFriendAddGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleNoticeFriendAdd(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleNoticeGroupRecall(data *model.EventNoticeGroupRecall, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleNoticeGroupRecall(ctx context.Context, in *EventNoticeGroupRecallGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleNoticeGroupRecall(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleNoticeFriendRecall(data *model.EventNoticeFriendRecall, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleNoticeFriendRecall(ctx context.Context, in *EventNoticeFriendRecallGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleNoticeFriendRecall(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleNoticeGroupNotifyPoke(data *model.EventNoticeGroupNotifyPoke, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleNoticeGroupNotifyPoke(ctx context.Context, in *EventNoticeGroupNotifyPokeGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleNoticeGroupNotifyPoke(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleNoticeGroupNotifyHonor(data *model.EventNoticeGroupNotifyHonor, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleNoticeGroupNotifyHonor(ctx context.Context, in *EventNoticeGroupNotifyHonorGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleNoticeGroupNotifyHonor(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleNoticeGroupNotifyLuckyKing(data *model.EventNoticeGroupNotifyLuckyKing, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleNoticeGroupNotifyLuckyKing(ctx context.Context, in *EventNoticeGroupNotifyLuckyKingGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleNoticeGroupNotifyLuckyKing(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleRequestFriend(data *model.EventRequestFriend, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleRequestFriend(ctx context.Context, in *EventRequestFriendGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleRequestFriend(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleRequestGroup(data *model.EventRequestGroup, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleRequestGroup(ctx context.Context, in *EventRequestGroupGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleRequestGroup(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleMetaLifecycle(data *model.EventMetaLifecycle, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleMetaLifecycle(ctx context.Context, in *EventMetaLifecycleGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleMetaLifecycle(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// HandleMetaHeartBeat(data *model.EventMetaHeartbeat, onebotApi api.OnebotApiClientInterface) error
func (svc *OnebotEventCallbackServerStub) HandleMetaHeartBeat(ctx context.Context, in *EventMetaHeartbeatGRPCWithOnebotApi) (*emptypb.Empty, error) {
	conn, err := svc.Broker.Dial(in.OnebotApi)
	if err != nil {
		log.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := svc.Impl.HandleMetaHeartBeat(in.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}
