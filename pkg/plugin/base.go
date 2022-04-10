package plugin

import (
	context "context"
	"log"
	"os/exec"

	"github.com/dezhishen/onebot-plus/pkg/cli"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// 插件的接口
type OnebotEventPlugin interface {
	//插件Id
	Id() string
	//插件名称
	Name() string
	//插件描述
	Description() string
	//插件帮助
	Help() string
	//=====插件的生命周期=====
	//初始化回调
	Init(cli.OnebotCli) error
	//退出前回调
	BeforeExit(cli.OnebotCli) error
	//=====Onebot事件回调=====
	//私聊消息
	MessagePrivate(*model.EventMessagePrivate, cli.OnebotCli) error
	//群组消息
	MessageGroup(*model.EventMessageGroup, cli.OnebotCli) error
	//生命周期
	MetaLifecycle(*model.EventMetaLifecycle, cli.OnebotCli) error
	//心跳
	MetaHeartbeat(*model.EventMetaHeartbeat, cli.OnebotCli) error
	//群文件上传
	NoticeGroupUpload(*model.EventNoticeGroupUpload, cli.OnebotCli) error
	//群管理员变动
	NoticeGroupAdmin(*model.EventNoticeGroupAdmin, cli.OnebotCli) error
	//群成员减少
	NoticeGroupDecrease(*model.EventNoticeGroupDecrease, cli.OnebotCli) error
	//群成员增加
	NoticeGroupIncrease(*model.EventNoticeGroupIncrease, cli.OnebotCli) error
	//群禁言
	NoticeGroupBan(*model.EventNoticeGroupBan, cli.OnebotCli) error
	//群消息撤回
	NoticeGroupRecall(*model.EventNoticeGroupRecall, cli.OnebotCli) error
	//群内戳一戳
	NoticeGroupNotifyPoke(*model.EventNoticeGroupNotifyPoke, cli.OnebotCli) error
	//群红包运气王
	NoticeGroupNotifyLuckyKing(*model.EventNoticeGroupNotifyLuckyKing, cli.OnebotCli) error
	//群成员荣誉变更
	NoticeGroupNotifyHonor(*model.EventNoticeGroupNotifyHonor, cli.OnebotCli) error
	//好友添加
	NoticeFriendAdd(*model.EventNoticeFriendAdd, cli.OnebotCli) error
	//好友消息撤回
	NoticeFriendRecall(*model.EventNoticeFriendRecall, cli.OnebotCli) error
	//好友添加请求
	RequestFriend(*model.EventRequestFriend, cli.OnebotCli) error
	//群添加/邀请请求
	RequestGroup(*model.EventRequestGroup, cli.OnebotCli) error
}

//插件实现
type onebotEventPluginGRPC struct {
	plugin.Plugin
	Impl OnebotEventPlugin
}

//插件实现GRPC的接口
func (p *onebotEventPluginGRPC) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterOnebotEventGRPCServer(s, &onebotEventPluginGRPCServerStub{Impl: p.Impl, broker: broker})
	return nil
}

type onebotEventPluginGRPCServerStub struct {
	// This is the real implementation
	broker *plugin.GRPCBroker
	Impl   OnebotEventPlugin
}

//插件Id
func (p *onebotEventPluginGRPCServerStub) Id(ctx context.Context, req *emptypb.Empty) (*wrapperspb.StringValue, error) {
	r := p.Impl.Id()
	return &wrapperspb.StringValue{Value: r}, nil
}

//插件名称
func (p *onebotEventPluginGRPCServerStub) Name(ctx context.Context, req *emptypb.Empty) (*wrapperspb.StringValue, error) {
	r := p.Impl.Name()
	return &wrapperspb.StringValue{Value: r}, nil
}

//插件描述
func (p *onebotEventPluginGRPCServerStub) Description(ctx context.Context, req *emptypb.Empty) (*wrapperspb.StringValue, error) {
	r := p.Impl.Description()
	return &wrapperspb.StringValue{Value: r}, nil
}

//插件帮助
func (p *onebotEventPluginGRPCServerStub) Help(ctx context.Context, req *emptypb.Empty) (*wrapperspb.StringValue, error) {
	r := p.Impl.Help()
	return &wrapperspb.StringValue{Value: r}, nil
}

//插件初始化
func (p *onebotEventPluginGRPCServerStub) Init(ctx context.Context, req *wrapperspb.UInt32Value) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Value)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	// defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.Init(client)
	return &emptypb.Empty{}, e
}

//插件退出前
func (p *onebotEventPluginGRPCServerStub) BeforeExit(ctx context.Context, req *wrapperspb.UInt32Value) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Value)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.BeforeExit(client)
	return &emptypb.Empty{}, e
}

//私聊消息
func (p *onebotEventPluginGRPCServerStub) MessagePrivate(ctx context.Context, req *EventMessagePrivateGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.MessagePrivate(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//群组消息
func (p *onebotEventPluginGRPCServerStub) MessageGroup(ctx context.Context, req *EventMessageGroupGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.MessageGroup(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//生命周期
func (p *onebotEventPluginGRPCServerStub) MetaLifecycle(ctx context.Context, req *EventMetaLifecycleGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.MetaLifecycle(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//心跳
func (p *onebotEventPluginGRPCServerStub) MetaHeartbeat(ctx context.Context, req *EventMetaHeartbeatGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.MetaHeartbeat(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//群文件上传
func (p *onebotEventPluginGRPCServerStub) NoticeGroupUpload(ctx context.Context, req *EventNoticeGroupUploadGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.NoticeGroupUpload(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//群管理员变化
func (p *onebotEventPluginGRPCServerStub) NoticeGroupAdmin(ctx context.Context, req *EventNoticeGroupAdminGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.NoticeGroupAdmin(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//群成员减少
func (p *onebotEventPluginGRPCServerStub) NoticeGroupDecrease(ctx context.Context, req *EventNoticeGroupDecreaseGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.NoticeGroupDecrease(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//群成员增加
func (p *onebotEventPluginGRPCServerStub) NoticeGroupIncrease(ctx context.Context, req *EventNoticeGroupIncreaseGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.NoticeGroupIncrease(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//群成员禁言
func (p *onebotEventPluginGRPCServerStub) NoticeGroupBan(ctx context.Context, req *EventNoticeGroupBanGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.NoticeGroupBan(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//群消息撤回
func (p *onebotEventPluginGRPCServerStub) NoticeGroupRecall(ctx context.Context, req *EventNoticeGroupRecallGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.NoticeGroupRecall(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//群内戳一戳
func (p *onebotEventPluginGRPCServerStub) NoticeGroupNotifyPoke(ctx context.Context, req *EventNoticeGroupNotifyPokeGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.NoticeGroupNotifyPoke(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//群红包运气王
func (p *onebotEventPluginGRPCServerStub) NoticeGroupNotifyLuckyKing(ctx context.Context, req *EventNoticeGroupNotifyLuckyKingGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.NoticeGroupNotifyLuckyKing(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//群成员荣誉变更
func (p *onebotEventPluginGRPCServerStub) NoticeGroupNotifyHonor(ctx context.Context, req *EventNoticeGroupNotifyHonorGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.NoticeGroupNotifyHonor(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//好友添加
func (p *onebotEventPluginGRPCServerStub) NoticeFriendAdd(ctx context.Context, req *EventNoticeFriendAddGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.NoticeFriendAdd(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//好友消息撤回
func (p *onebotEventPluginGRPCServerStub) NoticeFriendRecall(ctx context.Context, req *EventNoticeFriendRecallGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.NoticeFriendRecall(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//加好友请求
func (p *onebotEventPluginGRPCServerStub) RequestFriend(ctx context.Context, req *EventRequestFriendGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.RequestFriend(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//加群请求／邀请
func (p *onebotEventPluginGRPCServerStub) RequestGroup(ctx context.Context, req *EventRequestGroupGRPCWithCli) (*emptypb.Empty, error) {
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.OnebotCliClientStub{
		Client: cli.NewOnebotGrpcCliClient(conn),
	}
	e := p.Impl.RequestGroup(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

// 业务接口KV的实现，通过gRPC客户端转发请求给插件进程
type onebotEventPluginGRPCClientStub struct {
	broker *plugin.GRPCBroker
	client OnebotEventGRPCClient
}

//插件Id
func (m *onebotEventPluginGRPCClientStub) Id() string {
	// 转发
	grpc := &emptypb.Empty{}
	r, err := m.client.Id(context.Background(), grpc)
	if err != nil {
		logrus.Errorf("err:  %v", err)
		return ""
	}
	return r.Value
}

//插件名称
func (m *onebotEventPluginGRPCClientStub) Name() string {
	// 转发
	grpc := &emptypb.Empty{}
	r, err := m.client.Name(context.Background(), grpc)
	if err != nil {
		logrus.Errorf("err:  %v", err)
		return ""
	}
	return r.Value
}

//插件描述
func (m *onebotEventPluginGRPCClientStub) Description() string {
	// 转发
	grpc := &emptypb.Empty{}
	r, err := m.client.Description(context.Background(), grpc)
	if err != nil {
		logrus.Errorf("err:  %v", err)
		return ""
	}
	return r.Value
}

//插件帮助
func (m *onebotEventPluginGRPCClientStub) Help() string {
	// 转发
	grpc := &emptypb.Empty{}
	r, err := m.client.Help(context.Background(), grpc)
	if err != nil {
		logrus.Errorf("err:  %v", err)
		return ""
	}
	return r.Value
}

//初始化
func (m *onebotEventPluginGRPCClientStub) Init(onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	_, err := m.client.Init(context.Background(), &wrapperspb.UInt32Value{Value: brokerID})
	if err != nil {
		return err
	}
	// s.Stop()
	return err
}

//退出前回调
func (m *onebotEventPluginGRPCClientStub) BeforeExit(onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	_, err := m.client.BeforeExit(context.Background(), &wrapperspb.UInt32Value{Value: brokerID})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//私聊消息
func (m *onebotEventPluginGRPCClientStub) MessagePrivate(req *model.EventMessagePrivate, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.MessagePrivate(context.Background(), &EventMessagePrivateGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//群组消息
func (m *onebotEventPluginGRPCClientStub) MessageGroup(req *model.EventMessageGroup, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.MessageGroup(context.Background(), &EventMessageGroupGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//生命周期
func (m *onebotEventPluginGRPCClientStub) MetaLifecycle(req *model.EventMetaLifecycle, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.MetaLifecycle(context.Background(), &EventMetaLifecycleGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//心跳
func (m *onebotEventPluginGRPCClientStub) MetaHeartbeat(req *model.EventMetaHeartbeat, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.MetaHeartbeat(context.Background(), &EventMetaHeartbeatGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//群文件上传
func (m *onebotEventPluginGRPCClientStub) NoticeGroupUpload(req *model.EventNoticeGroupUpload, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.NoticeGroupUpload(context.Background(), &EventNoticeGroupUploadGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//群管理员变动
func (m *onebotEventPluginGRPCClientStub) NoticeGroupAdmin(req *model.EventNoticeGroupAdmin, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.NoticeGroupAdmin(context.Background(), &EventNoticeGroupAdminGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//群成员减少
func (m *onebotEventPluginGRPCClientStub) NoticeGroupDecrease(req *model.EventNoticeGroupDecrease, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.NoticeGroupDecrease(context.Background(), &EventNoticeGroupDecreaseGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//群成员增加
func (m *onebotEventPluginGRPCClientStub) NoticeGroupIncrease(req *model.EventNoticeGroupIncrease, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.NoticeGroupIncrease(context.Background(), &EventNoticeGroupIncreaseGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//群禁言
func (m *onebotEventPluginGRPCClientStub) NoticeGroupBan(req *model.EventNoticeGroupBan, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.NoticeGroupBan(context.Background(), &EventNoticeGroupBanGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//群消息撤回
func (m *onebotEventPluginGRPCClientStub) NoticeGroupRecall(req *model.EventNoticeGroupRecall, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.NoticeGroupRecall(context.Background(), &EventNoticeGroupRecallGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//群内戳一戳
func (m *onebotEventPluginGRPCClientStub) NoticeGroupNotifyPoke(req *model.EventNoticeGroupNotifyPoke, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.NoticeGroupNotifyPoke(context.Background(), &EventNoticeGroupNotifyPokeGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//群红包运气王
func (m *onebotEventPluginGRPCClientStub) NoticeGroupNotifyLuckyKing(req *model.EventNoticeGroupNotifyLuckyKing, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.NoticeGroupNotifyLuckyKing(context.Background(), &EventNoticeGroupNotifyLuckyKingGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//群成员荣誉变更
func (m *onebotEventPluginGRPCClientStub) NoticeGroupNotifyHonor(req *model.EventNoticeGroupNotifyHonor, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.NoticeGroupNotifyHonor(context.Background(), &EventNoticeGroupNotifyHonorGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//好友添加
func (m *onebotEventPluginGRPCClientStub) NoticeFriendAdd(req *model.EventNoticeFriendAdd, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.NoticeFriendAdd(context.Background(), &EventNoticeFriendAddGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//好友消息撤回
func (m *onebotEventPluginGRPCClientStub) NoticeFriendRecall(req *model.EventNoticeFriendRecall, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.NoticeFriendRecall(context.Background(), &EventNoticeFriendRecallGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//好友添加请求
func (m *onebotEventPluginGRPCClientStub) RequestFriend(req *model.EventRequestFriend, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.RequestFriend(context.Background(), &EventRequestFriendGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//群添加/邀请请求
func (m *onebotEventPluginGRPCClientStub) RequestGroup(req *model.EventRequestGroup, onebotCli cli.OnebotCli) error {
	// 转发
	messageCliServer := &cli.OnebotCliServerStub{
		Impl: onebotCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterOnebotGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.RequestGroup(context.Background(), &EventRequestGroupGRPCWithCli{
		Message: msg,
		Cli:     brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return err
}

//插件实现GRPC的接口
func (p *onebotEventPluginGRPC) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &onebotEventPluginGRPCClientStub{
		broker: broker,
		client: NewOnebotEventGRPCClient(c)}, nil
}

func LoadPlugin(path string) (OnebotEventPlugin, *plugin.Client) {
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"main": &onebotEventPluginGRPC{},
		},
		Cmd: exec.Command(path),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	// defer client.Kill()
	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		logrus.Error(err)
	}
	// Request the plugin
	raw, err := rpcClient.Dispense("main")
	if err != nil {
		log.Fatal(err)
	}
	d := raw.(OnebotEventPlugin)
	return d, client
}
