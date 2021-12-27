package event

import (
	context "context"

	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type EventPrivateMessageHandler interface {
	//处理消息
	Handler(model.EventMessagePrivate) error
}

type EventPrivateMessageHandlerGRPCPlugin struct {
	// 需要嵌入插件接口
	plugin.Plugin
	// 具体实现，仅当业务接口实现基于Go时该字段有用
	Impl EventPrivateMessageHandler
}

//插件实现GRPC的接口
func (p *EventPrivateMessageHandlerGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterEventPrivateMsgServer(s, &EventPrivateMessageHandlerServerStub{Impl: p.Impl})
	return nil
}

type EventPrivateMessageHandlerServerStub struct {
	// This is the real implementation
	Impl EventPrivateMessageHandler
}

//处理消息
func (m *EventPrivateMessageHandlerServerStub) Handle(ctx context.Context, in *model.EventMessagePrivateGRPC) (*emptypb.Empty, error) {
	e := m.Impl.Handler(*in.ToStruct())
	return &emptypb.Empty{}, e
}

// 业务接口的实现，通过gRPC客户端转发请求给插件进程
type EventPrivateMessageHandlerClientStub struct {
	client EventPrivateMsgClient
}

//处理
func (m *EventPrivateMessageHandlerClientStub) Handle(msg *model.EventMessagePrivate) error {
	// 转发
	_, err := m.client.Handle(context.Background(), msg.ToGRPC())
	return err
}

//插件实现GRPC的接口
func (p *EventPrivateMessageHandlerGRPCPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &EventPrivateMessageHandlerClientStub{client: NewEventPrivateMsgClient(c)}, nil
}
