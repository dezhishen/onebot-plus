package plugin

import (
	"context"
	"log"

	"github.com/dezhishen/onebot-plus/pkg/event"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

//插件接口
type PrivateMssageHanlder interface {
	Handle(*model.EventMessagePrivate) error
}

//插件实现
type PrivateMessageGRPCPlugin struct {
	plugin.Plugin
	Impl PrivateMssageHanlder
}

//插件实现GRPC的接口
func (p *PrivateMessageGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	event.RegisterEventPrivateMsgServer(s, &PrivateMessageGRPCServerStub{Impl: p.Impl})
	return nil
}

//声明插件类型
func (p *PrivateMessageGRPCPlugin) EventType() string {
	return "private_message"
}

type PrivateMessageGRPCServerStub struct {
	// This is the real implementation
	Impl PrivateMssageHanlder
}

func (m *PrivateMessageGRPCServerStub) Handle(ctx context.Context, req *model.EventMessagePrivateGRPC) (*emptypb.Empty, error) {
	log.Printf("GRPC调用...%v", req)
	e := m.Impl.Handle(req.ToStruct())
	return &emptypb.Empty{}, e
}

// 业务接口KV的实现，通过gRPC客户端转发请求给插件进程
type PrivateMessageGRPCClientStub struct {
	client event.EventPrivateMsgClient
}

//业务接口
func (m *PrivateMessageGRPCClientStub) Handle(req *model.EventMessagePrivate) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.Handle(context.Background(), grpc)
	if err != nil {
		logrus.Errorf("err:  %v", err)
	}
	return err
}

//插件实现GRPC的接口
func (p *PrivateMessageGRPCPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &PrivateMessageGRPCClientStub{client: event.NewEventPrivateMsgClient(c)}, nil
}
