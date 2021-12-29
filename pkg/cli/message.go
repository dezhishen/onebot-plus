package cli

import (
	context "context"

	"github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type MessageCli interface {
	//发送消息
	SendMsg(msg *model.MsgForSend) (int64, error)
	//发送私聊消息
	SendPrivateMsg(msg *model.PrivateMsg) (int64, error)
	// 发送群消息
	SendGroupMsg(msg *model.GroupMsg) (int64, error)
	//删除消息
	DelMsg(id int64) error
	//获取消息
	GetMsg(id int64) (*model.MessageData, error)
	//获取转发的消息
	GetForwardMsg(id int64) (*model.ForwardMessageData, error)
}

type MessageCliGRPCPlugin struct {
	// 需要嵌入插件接口
	plugin.Plugin
	// 具体实现，仅当业务接口实现基于Go时该字段有用
	Impl MessageCli
}

//插件实现GRPC的接口
func (p *MessageCliGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterMessageGrpcCliServer(s, &MessageCliServerStub{Impl: p.Impl})
	return nil
}

type MessageCliServerStub struct {
	// This is the real implementation
	Impl MessageCli
}

//删除消息
func (m *MessageCliServerStub) DelMsg(ctx context.Context, in *wrapperspb.Int64Value) (*emptypb.Empty, error) {
	e := m.Impl.DelMsg(in.Value)
	return &emptypb.Empty{}, e
}

//发送消息
func (m *MessageCliServerStub) SendMsg(ctx context.Context, in *model.MsgForSendGRPC) (*wrapperspb.Int64Value, error) {
	v, e := m.Impl.SendMsg(in.ToStruct())
	return &wrapperspb.Int64Value{Value: v}, e
}

//发送私聊消息
func (m *MessageCliServerStub) SendPrivateMsg(ctx context.Context, in *model.PrivateMsgGRPC) (*wrapperspb.Int64Value, error) {
	v, e := m.Impl.SendPrivateMsg(in.ToStruct())
	return &wrapperspb.Int64Value{Value: v}, e
}

// 发送群消息
func (m *MessageCliServerStub) SendGroupMsg(ctx context.Context, in *model.GroupMsgGRPC) (*wrapperspb.Int64Value, error) {
	v, e := m.Impl.SendGroupMsg(in.ToStruct())
	return &wrapperspb.Int64Value{Value: v}, e
}

//获取消息
func (m *MessageCliServerStub) GetMsg(ctx context.Context, in *wrapperspb.Int64Value) (*model.MessageDataGRPC, error) {
	v, e := m.Impl.GetMsg(in.Value)
	return v.ToGRPC(), e
}

//获取转发的消息
func (m *MessageCliServerStub) GetForwardMsg(ctx context.Context, in *wrapperspb.Int64Value) (*model.ForwardMessageDataGRPC, error) {
	v, e := m.Impl.GetForwardMsg(in.Value)
	return v.ToGRPC(), e
}

// 业务接口的实现，通过gRPC客户端转发请求给插件进程
type MessageCliClientStub struct {
	Client MessageGrpcCliClient
}

func NewMessageCliClientStub(cli MessageGrpcCliClient) *MessageCliClientStub {
	return &MessageCliClientStub{
		Client: cli,
	}
}

//发送消息
func (m *MessageCliClientStub) SendMsg(msg *model.MsgForSend) (int64, error) {
	// 转发
	resp, err := m.Client.SendMsg(context.Background(), msg.ToGRPC())
	if err != nil {
		return 0, err
	}
	return resp.Value, err
}

//发送私聊消息
func (m *MessageCliClientStub) SendPrivateMsg(msg *model.PrivateMsg) (int64, error) {
	// 转发
	resp, err := m.Client.SendPrivateMsg(context.Background(), msg.ToGRPC())
	if err != nil {
		return 0, err
	}
	return resp.Value, err
}

//发送群组消息
func (m *MessageCliClientStub) SendGroupMsg(msg *model.GroupMsg) (int64, error) {
	// 转发
	resp, err := m.Client.SendGroupMsg(context.Background(), msg.ToGRPC())
	if err != nil {
		return 0, err
	}
	return resp.Value, err
}

//删除消息
func (m *MessageCliClientStub) DelMsg(id int64) error {
	// 转发
	_, err := m.Client.DelMsg(context.Background(), &wrapperspb.Int64Value{Value: id})
	return err
}

//获取消息
func (m *MessageCliClientStub) GetMsg(id int64) (*model.MessageData, error) {
	// 转发
	logrus.Infof("获取消息,ID:[%v]", id)
	resp, err := m.Client.GetMsg(context.Background(), &wrapperspb.Int64Value{Value: id})
	if err != nil {
		return nil, err
	}
	return resp.ToStruct(), err
}

//获取转发的消息
func (m *MessageCliClientStub) GetForwardMsg(id int64) (*model.ForwardMessageData, error) {
	// 转发
	resp, err := m.Client.GetForwardMsg(context.Background(), &wrapperspb.Int64Value{Value: id})
	if err != nil {
		return nil, err
	}
	return resp.ToStruct(), err
}

//插件实现GRPC的接口
func (p *MessageCliGRPCPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &MessageCliClientStub{Client: NewMessageGrpcCliClient(c)}, nil
}
