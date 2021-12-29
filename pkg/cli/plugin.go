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

type OnebotCliGRPCPlugin struct {
	// 需要嵌入插件接口
	plugin.Plugin
	// 具体实现，仅当业务接口实现基于Go时该字段有用
	Impl OnebotCli
}

//插件实现GRPC的接口
func (p *OnebotCliGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterOnebotGrpcCliServer(s, &OnebotCliServerStub{Impl: p.Impl})
	return nil
}

type OnebotCliServerStub struct {
	// This is the real implementation
	Impl OnebotCli
}

//删除消息
func (m *OnebotCliServerStub) DelMsg(ctx context.Context, in *wrapperspb.Int64Value) (*emptypb.Empty, error) {
	e := m.Impl.DelMsg(in.Value)
	return &emptypb.Empty{}, e
}

//发送消息
func (m *OnebotCliServerStub) SendMsg(ctx context.Context, in *model.MsgForSendGRPC) (*wrapperspb.Int64Value, error) {
	v, e := m.Impl.SendMsg(in.ToStruct())
	return &wrapperspb.Int64Value{Value: v}, e
}

//发送私聊消息
func (m *OnebotCliServerStub) SendPrivateMsg(ctx context.Context, in *model.PrivateMsgGRPC) (*wrapperspb.Int64Value, error) {
	v, e := m.Impl.SendPrivateMsg(in.ToStruct())
	return &wrapperspb.Int64Value{Value: v}, e
}

// 发送群消息
func (m *OnebotCliServerStub) SendGroupMsg(ctx context.Context, in *model.GroupMsgGRPC) (*wrapperspb.Int64Value, error) {
	v, e := m.Impl.SendGroupMsg(in.ToStruct())
	return &wrapperspb.Int64Value{Value: v}, e
}

//获取消息
func (m *OnebotCliServerStub) GetMsg(ctx context.Context, in *wrapperspb.Int64Value) (*model.MessageDataGRPC, error) {
	v, e := m.Impl.GetMsg(in.Value)
	return v.ToGRPC(), e
}

//获取转发的消息
func (m *OnebotCliServerStub) GetForwardMsg(ctx context.Context, in *wrapperspb.Int64Value) (*model.ForwardMessageDataGRPC, error) {
	v, e := m.Impl.GetForwardMsg(in.Value)
	return v.ToGRPC(), e
}

//获取登录信息
func (m *OnebotCliServerStub) GetLoginInfo(ctx context.Context, in *emptypb.Empty) (*model.AccountGRPC, error) {
	v, e := m.Impl.GetLoginInfo()
	return v.ToGRPC(), e
}

//获取陌生人信息
func (m *OnebotCliServerStub) GetStrangerInfo(ctx context.Context, in *GetStrangerInfoReq) (*model.AccountGRPC, error) {
	v, e := m.Impl.GetLoginInfo()
	return v.ToGRPC(), e
}
func (m *OnebotCliServerStub) GetCookies(ctx context.Context, in *wrapperspb.StringValue) (*model.CokiesGRPC, error) {
	v, e := m.Impl.GetCookies(in.Value)
	return v.ToGRPC(), e
}
func (m *OnebotCliServerStub) GetCSRFToken(ctx context.Context, in *emptypb.Empty) (*model.CSRFTokenGRPC, error) {
	v, e := m.Impl.GetCSRFToken()
	return v.ToGRPC(), e
}
func (m *OnebotCliServerStub) GetCredentials(ctx context.Context, in *wrapperspb.StringValue) (*model.CredentialsGRPC, error) {
	v, e := m.Impl.GetCredentials(in.Value)
	return v.ToGRPC(), e
}

//获取语音
func (m *OnebotCliServerStub) GetRecord(ctx context.Context, in *GetRecordReq) (*model.FileGRPC, error) {
	v, e := m.Impl.GetRecord(in.File, in.OutFormat)
	return v.ToGRPC(), e
}

//获取图片
func (m *OnebotCliServerStub) GetImage(ctx context.Context, in *wrapperspb.StringValue) (*model.FileGRPC, error) {
	v, e := m.Impl.GetImage(in.Value)
	return v.ToGRPC(), e
}

// 业务接口的实现，通过gRPC客户端转发请求给插件进程
type OnebotCliClientStub struct {
	Client OnebotGrpcCliClient
}

//发送消息
func (m *OnebotCliClientStub) SendMsg(msg *model.MsgForSend) (int64, error) {
	// 转发
	resp, err := m.Client.SendMsg(context.Background(), msg.ToGRPC())
	if err != nil {
		return 0, err
	}
	return resp.Value, err
}

//发送私聊消息
func (m *OnebotCliClientStub) SendPrivateMsg(msg *model.PrivateMsg) (int64, error) {
	// 转发
	resp, err := m.Client.SendPrivateMsg(context.Background(), msg.ToGRPC())
	if err != nil {
		return 0, err
	}
	return resp.Value, err
}

//发送群组消息
func (m *OnebotCliClientStub) SendGroupMsg(msg *model.GroupMsg) (int64, error) {
	// 转发
	resp, err := m.Client.SendGroupMsg(context.Background(), msg.ToGRPC())
	if err != nil {
		return 0, err
	}
	return resp.Value, err
}

//删除消息
func (m *OnebotCliClientStub) DelMsg(id int64) error {
	// 转发
	_, err := m.Client.DelMsg(context.Background(), &wrapperspb.Int64Value{Value: id})
	return err
}

//获取消息
func (m *OnebotCliClientStub) GetMsg(id int64) (*model.MessageData, error) {
	// 转发
	logrus.Infof("获取消息,ID:[%v]", id)
	resp, err := m.Client.GetMsg(context.Background(), &wrapperspb.Int64Value{Value: id})
	if err != nil {
		return nil, err
	}
	return resp.ToStruct(), err
}

//获取转发的消息
func (m *OnebotCliClientStub) GetForwardMsg(id int64) (*model.ForwardMessageData, error) {
	// 转发
	resp, err := m.Client.GetForwardMsg(context.Background(), &wrapperspb.Int64Value{Value: id})
	if err != nil {
		return nil, err
	}
	return resp.ToStruct(), err
}

//获取登录信息
func (m *OnebotCliClientStub) GetLoginInfo() (*model.Account, error) {
	// 转发
	resp, err := m.Client.GetLoginInfo(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return resp.ToStruct(), err
}

//获取陌生人信息
func (m *OnebotCliClientStub) GetStrangerInfo(userId int64, noCache bool) (*model.Account, error) {
	// 转发
	resp, err := m.Client.GetStrangerInfo(context.Background(), &GetStrangerInfoReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return resp.ToStruct(), err
}

func (m *OnebotCliClientStub) GetCookies(domin string) (*model.Cokies, error) {
	// 转发
	resp, err := m.Client.GetCookies(context.Background(), &wrapperspb.StringValue{
		Value: domin,
	})
	if err != nil {
		return nil, err
	}
	return resp.ToStruct(), err
}

func (m *OnebotCliClientStub) GetCSRFToken() (*model.CSRFToken, error) {
	// 转发
	resp, err := m.Client.GetCSRFToken(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return resp.ToStruct(), err
}

func (m *OnebotCliClientStub) GetCredentials(domin string) (*model.Credentials, error) {
	// 转发
	resp, err := m.Client.GetCredentials(context.Background(), &wrapperspb.StringValue{Value: domin})
	if err != nil {
		return nil, err
	}
	return resp.ToStruct(), err
}

//获取语音
func (m *OnebotCliClientStub) GetRecord(file string, out_format string) (*model.File, error) {
	// 转发
	resp, err := m.Client.GetRecord(context.Background(), &GetRecordReq{File: file, OutFormat: out_format})
	if err != nil {
		return nil, err
	}
	return resp.ToStruct(), err
}

//获取图片
func (m *OnebotCliClientStub) GetImage(file string) (*model.File, error) {
	// 转发
	resp, err := m.Client.GetImage(context.Background(), &wrapperspb.StringValue{Value: file})
	if err != nil {
		return nil, err
	}
	return resp.ToStruct(), err
}

//插件实现GRPC的接口
func (p *OnebotCliGRPCPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &OnebotCliClientStub{Client: NewOnebotGrpcCliClient(c)}, nil
}
