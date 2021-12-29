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
	//私聊消息
	MessagePrivate(*model.EventMessagePrivate, cli.MessageCli) error
	//群组消息
	MessageGroup(*model.EventMessageGroup) error
	//生命周期
	MetaLifecycle(*model.EventMetaLifecycle) error
	//心跳
	MetaHeartbeat(*model.EventMetaHeartbeat) error
	//群文件上传
	NoticeGroupUpload(*model.EventNoticeGroupUpload) error
	//群管理员变动
	NoticeGroupAdmin(*model.EventNoticeGroupAdmin) error
	//群成员减少
	NoticeGroupDecrease(*model.EventNoticeGroupDecrease) error
	//群成员增加
	NoticeGroupIncrease(*model.EventNoticeGroupIncrease) error
	//群禁言
	NoticeGroupBan(*model.EventNoticeGroupBan) error
	//群消息撤回
	NoticeGroupRecall(*model.EventNoticeGroupRecall) error
	//群内戳一戳
	NoticeGroupNotifyPoke(*model.EventNoticeGroupNotifyPoke) error
	//群红包运气王
	NoticeGroupNotifyLuckyKing(*model.EventNoticeGroupNotifyLuckyKing) error
	//群成员荣誉变更
	NoticeGroupNotifyHonor(*model.EventNoticeGroupNotifyHonor) error
	//好友添加
	NoticeFriendAdd(*model.EventNoticeFriendAdd) error
	//好友消息撤回
	NoticeFriendRecall(*model.EventNoticeFriendRecall) error
	//好友添加请求
	RequestFriend(*model.EventRequestFriend) error
	//群添加/邀请请求
	RequestGroup(*model.EventRequestGroup) error
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

//私聊消息
func (p *onebotEventPluginGRPCServerStub) MessagePrivate(ctx context.Context, req *EventMessagePrivateGRPCWithCliGRPC) (*emptypb.Empty, error) {
	logrus.Infof("req....%v", req)
	logrus.Info("req.Cli...", req.Cli)
	conn, err := p.broker.Dial(req.Cli)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &cli.MessageCliClientStub{
		Client: cli.NewMessageGrpcCliClient(conn),
	}
	logrus.Infof("创建客户端....%v", client)
	e := p.Impl.MessagePrivate(req.Message.ToStruct(), client)
	return &emptypb.Empty{}, e
}

//群组消息
func (p *onebotEventPluginGRPCServerStub) MessageGroup(ctx context.Context, req *model.EventMessageGroupGRPC) (*emptypb.Empty, error) {
	e := p.Impl.MessageGroup(req.ToStruct())
	return &emptypb.Empty{}, e
}

//生命周期
func (p *onebotEventPluginGRPCServerStub) MetaLifecycle(ctx context.Context, req *model.EventMetaLifecycleGRPC) (*emptypb.Empty, error) {
	e := p.Impl.MetaLifecycle(req.ToStruct())
	return &emptypb.Empty{}, e
}

//心跳
func (p *onebotEventPluginGRPCServerStub) MetaHeartbeat(ctx context.Context, req *model.EventMetaHeartbeatGRPC) (*emptypb.Empty, error) {
	e := p.Impl.MetaHeartbeat(req.ToStruct())
	return &emptypb.Empty{}, e
}

//群文件上传
func (p *onebotEventPluginGRPCServerStub) NoticeGroupUpload(ctx context.Context, req *model.EventNoticeGroupUploadGRPC) (*emptypb.Empty, error) {
	e := p.Impl.NoticeGroupUpload(req.ToStruct())
	return &emptypb.Empty{}, e
}

//群管理员变化
func (p *onebotEventPluginGRPCServerStub) NoticeGroupAdmin(ctx context.Context, req *model.EventNoticeGroupAdminGRPC) (*emptypb.Empty, error) {
	e := p.Impl.NoticeGroupAdmin(req.ToStruct())
	return &emptypb.Empty{}, e
}

//群成员减少
func (p *onebotEventPluginGRPCServerStub) NoticeGroupDecrease(ctx context.Context, req *model.EventNoticeGroupDecreaseGRPC) (*emptypb.Empty, error) {
	e := p.Impl.NoticeGroupDecrease(req.ToStruct())
	return &emptypb.Empty{}, e
}

//群成员增加
func (p *onebotEventPluginGRPCServerStub) NoticeGroupIncrease(ctx context.Context, req *model.EventNoticeGroupIncreaseGRPC) (*emptypb.Empty, error) {
	e := p.Impl.NoticeGroupIncrease(req.ToStruct())
	return &emptypb.Empty{}, e
}

//群成员禁言
func (p *onebotEventPluginGRPCServerStub) NoticeGroupBan(ctx context.Context, req *model.EventNoticeGroupBanGRPC) (*emptypb.Empty, error) {
	e := p.Impl.NoticeGroupBan(req.ToStruct())
	return &emptypb.Empty{}, e
}

//群消息撤回
func (p *onebotEventPluginGRPCServerStub) NoticeGroupRecall(ctx context.Context, req *model.EventNoticeGroupRecallGRPC) (*emptypb.Empty, error) {
	e := p.Impl.NoticeGroupRecall(req.ToStruct())
	return &emptypb.Empty{}, e
}

//群内戳一戳
func (p *onebotEventPluginGRPCServerStub) NoticeGroupNotifyPoke(ctx context.Context, req *model.EventNoticeGroupNotifyPokeGRPC) (*emptypb.Empty, error) {
	e := p.Impl.NoticeGroupNotifyPoke(req.ToStruct())
	return &emptypb.Empty{}, e
}

//群红包运气王
func (p *onebotEventPluginGRPCServerStub) NoticeGroupNotifyLuckyKing(ctx context.Context, req *model.EventNoticeGroupNotifyLuckyKingGRPC) (*emptypb.Empty, error) {
	e := p.Impl.NoticeGroupNotifyLuckyKing(req.ToStruct())
	return &emptypb.Empty{}, e
}

//群成员荣誉变更
func (p *onebotEventPluginGRPCServerStub) NoticeGroupNotifyHonor(ctx context.Context, req *model.EventNoticeGroupNotifyHonorGRPC) (*emptypb.Empty, error) {
	e := p.Impl.NoticeGroupNotifyHonor(req.ToStruct())
	return &emptypb.Empty{}, e
}

//好友添加
func (p *onebotEventPluginGRPCServerStub) NoticeFriendAdd(ctx context.Context, req *model.EventNoticeFriendAddGRPC) (*emptypb.Empty, error) {
	e := p.Impl.NoticeFriendAdd(req.ToStruct())
	return &emptypb.Empty{}, e
}

//好友消息撤回
func (p *onebotEventPluginGRPCServerStub) NoticeFriendRecall(ctx context.Context, req *model.EventNoticeFriendRecallGRPC) (*emptypb.Empty, error) {
	e := p.Impl.NoticeFriendRecall(req.ToStruct())
	return &emptypb.Empty{}, e
}

//加好友请求
func (p *onebotEventPluginGRPCServerStub) RequestFriend(ctx context.Context, req *model.EventRequestFriendGRPC) (*emptypb.Empty, error) {
	e := p.Impl.RequestFriend(req.ToStruct())
	return &emptypb.Empty{}, e
}

//加群请求／邀请
func (p *onebotEventPluginGRPCServerStub) RequestGroup(ctx context.Context, req *model.EventRequestGroupGRPC) (*emptypb.Empty, error) {
	e := p.Impl.RequestGroup(req.ToStruct())
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

//私聊消息
func (m *onebotEventPluginGRPCClientStub) MessagePrivate(req *model.EventMessagePrivate, msgCli cli.MessageCli) error {
	// 转发
	messageCliServer := &cli.MessageCliServerStub{
		Impl: msgCli,
	} //{Impl: cli}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		cli.RegisterMessageGrpcCliServer(s, messageCliServer)
		return s
	}
	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	msg := req.ToGRPC()
	_, err := m.client.MessagePrivate(context.Background(), &EventMessagePrivateGRPCWithCliGRPC{
		Message: msg,
		Cli:     brokerID,
	})
	s.Stop()
	return err
}

//群组消息
func (m *onebotEventPluginGRPCClientStub) MessageGroup(req *model.EventMessageGroup) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.MessageGroup(context.Background(), grpc)
	return err
}

//生命周期
func (m *onebotEventPluginGRPCClientStub) MetaLifecycle(req *model.EventMetaLifecycle) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.MetaLifecycle(context.Background(), grpc)
	return err
}

//心跳
func (m *onebotEventPluginGRPCClientStub) MetaHeartbeat(req *model.EventMetaHeartbeat) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.MetaHeartbeat(context.Background(), grpc)
	return err
}

//群文件上传
func (m *onebotEventPluginGRPCClientStub) NoticeGroupUpload(req *model.EventNoticeGroupUpload) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.NoticeGroupUpload(context.Background(), grpc)
	return err
}

//群管理员变动
func (m *onebotEventPluginGRPCClientStub) NoticeGroupAdmin(req *model.EventNoticeGroupAdmin) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.NoticeGroupAdmin(context.Background(), grpc)
	return err
}

//群成员减少
func (m *onebotEventPluginGRPCClientStub) NoticeGroupDecrease(req *model.EventNoticeGroupDecrease) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.NoticeGroupDecrease(context.Background(), grpc)
	return err
}

//群成员增加
func (m *onebotEventPluginGRPCClientStub) NoticeGroupIncrease(req *model.EventNoticeGroupIncrease) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.NoticeGroupIncrease(context.Background(), grpc)
	return err
}

//群禁言
func (m *onebotEventPluginGRPCClientStub) NoticeGroupBan(req *model.EventNoticeGroupBan) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.NoticeGroupBan(context.Background(), grpc)
	return err
}

//群消息撤回
func (m *onebotEventPluginGRPCClientStub) NoticeGroupRecall(req *model.EventNoticeGroupRecall) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.NoticeGroupRecall(context.Background(), grpc)
	return err
}

//群内戳一戳
func (m *onebotEventPluginGRPCClientStub) NoticeGroupNotifyPoke(req *model.EventNoticeGroupNotifyPoke) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.NoticeGroupNotifyPoke(context.Background(), grpc)
	return err
}

//群红包运气王
func (m *onebotEventPluginGRPCClientStub) NoticeGroupNotifyLuckyKing(req *model.EventNoticeGroupNotifyLuckyKing) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.NoticeGroupNotifyLuckyKing(context.Background(), grpc)
	return err
}

//群成员荣誉变更
func (m *onebotEventPluginGRPCClientStub) NoticeGroupNotifyHonor(req *model.EventNoticeGroupNotifyHonor) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.NoticeGroupNotifyHonor(context.Background(), grpc)
	return err
}

//好友添加
func (m *onebotEventPluginGRPCClientStub) NoticeFriendAdd(req *model.EventNoticeFriendAdd) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.NoticeFriendAdd(context.Background(), grpc)
	return err
}

//好友消息撤回
func (m *onebotEventPluginGRPCClientStub) NoticeFriendRecall(req *model.EventNoticeFriendRecall) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.NoticeFriendRecall(context.Background(), grpc)
	return err
}

//好友添加请求
func (m *onebotEventPluginGRPCClientStub) RequestFriend(req *model.EventRequestFriend) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.RequestFriend(context.Background(), grpc)
	return err
}

//群添加/邀请请求
func (m *onebotEventPluginGRPCClientStub) RequestGroup(req *model.EventRequestGroup) error {
	// 转发
	grpc := req.ToGRPC()
	_, err := m.client.RequestGroup(context.Background(), grpc)
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
