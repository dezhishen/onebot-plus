package event

import (
	context "context"

	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type EventGroupMessageHandler interface {
	//处理消息
	Handler(model.EventMessageGroup) error
}

type EventGroupMessageHandlerGRPCPlugin struct {
	// 需要嵌入插件接口
	plugin.Plugin
	// 具体实现，仅当业务接口实现基于Go时该字段有用
	Impl EventGroupMessageHandler
}

//插件实现GRPC的接口
func (p *EventGroupMessageHandlerGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterEventGroupMsgServer(s, &EventGroupMessageHandlerServerStub{Impl: p.Impl})
	return nil
}

type EventGroupMessageHandlerServerStub struct {
	// This is the real implementation
	Impl EventGroupMessageHandler
}

//处理消息
func (m *EventGroupMessageHandlerServerStub) Handle(ctx context.Context, in *model.EventMessageGroupGRPC) (*emptypb.Empty, error) {
	e := m.Impl.Handler(*in.ToStruct())
	return &emptypb.Empty{}, e
}

// 业务接口的实现，通过gRPC客户端转发请求给插件进程
type EventGroupMessageHandlerClientStub struct {
	client EventGroupMsgClient
}

//处理
func (m *EventGroupMessageHandlerClientStub) Handle(msg *model.EventMessageGroup) error {
	// 转发
	_, err := m.client.Handle(context.Background(), msg.ToGRPC())
	return err
}

//插件实现GRPC的接口
func (p *EventGroupMessageHandlerGRPCPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &EventGroupMessageHandlerClientStub{client: NewEventGroupMsgClient(c)}, nil
}
