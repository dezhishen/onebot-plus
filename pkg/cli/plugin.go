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
func (m *OnebotCliServerStub) SendLike(ctx context.Context, in *SendLikeReq) (*emptypb.Empty, error) {
	e := m.Impl.SendLike(in.UserId, in.Times)
	return &emptypb.Empty{}, e
}

//处理加好友请求
func (m *OnebotCliServerStub) SetFriendAddRequest(ctx context.Context, in *SetFriendAddRequestReq) (*emptypb.Empty, error) {
	e := m.Impl.SetFriendAddRequest(in.Flag, in.Approve, in.Remark)
	return &emptypb.Empty{}, e
}
func (m *OnebotCliServerStub) GetFriendList(ctx context.Context, in *emptypb.Empty) (*GetFriendListResp, error) {
	v, e := m.Impl.GetFriendList()
	return &GetFriendListResp{Data: model.AccountArray2AccountGRPCArray(v)}, e
}

// 群组踢人
func (m *OnebotCliServerStub) SetGroupKick(ctx context.Context, in *SetGroupKickReq) (*emptypb.Empty, error) {
	e := m.Impl.SetGroupKick(in.GroupId, in.UserId, in.RejectRddRequest)
	return &emptypb.Empty{}, e
}

// 群组禁言
func (m *OnebotCliServerStub) SetGroupBan(ctx context.Context, in *SetGroupBanReq) (*emptypb.Empty, error) {
	e := m.Impl.SetGroupBan(in.GroupId, in.UserId, in.Duration)
	return &emptypb.Empty{}, e
}

// 群组匿名用户禁言
func (m *OnebotCliServerStub) SetGroupAnonymousBan(ctx context.Context, in *SetGroupAnonymousBanReq) (*emptypb.Empty, error) {
	e := m.Impl.SetGroupAnonymousBan(in.GroupId, in.Duration, in.AnonymousFlag)
	return &emptypb.Empty{}, e
}

//群组全员禁言
func (m *OnebotCliServerStub) SetGroupWholeBan(ctx context.Context, in *SetGroupWholeBanReq) (*emptypb.Empty, error) {
	e := m.Impl.SetGroupWholeBan(in.GroupId, in.Enable)
	return &emptypb.Empty{}, e
}

//群组设置管理员
func (m *OnebotCliServerStub) SetGroupAdmin(ctx context.Context, in *SetGroupAdminReq) (*emptypb.Empty, error) {
	e := m.Impl.SetGroupAdmin(in.GroupId, in.UserId, in.Enable)
	return &emptypb.Empty{}, e
}

//群组匿名
func (m *OnebotCliServerStub) SetGroupAnonymous(ctx context.Context, in *SetGroupAnonymousReq) (*emptypb.Empty, error) {
	e := m.Impl.SetGroupAnonymous(in.GroupId, in.Enable)
	return &emptypb.Empty{}, e
}

//设置群名片（群备注）
func (m *OnebotCliServerStub) SetGroupCard(ctx context.Context, in *SetGroupCardReq) (*emptypb.Empty, error) {
	e := m.Impl.SetGroupCard(in.GroupId, in.UserId, in.Card)
	return &emptypb.Empty{}, e
}

//设置群名
func (m *OnebotCliServerStub) SetGroupName(ctx context.Context, in *SetGroupNameReq) (*emptypb.Empty, error) {
	e := m.Impl.SetGroupName(in.GroupId, in.GroupName)
	return &emptypb.Empty{}, e
}

//退出群组
func (m *OnebotCliServerStub) SetGroupLeave(ctx context.Context, in *SetGroupLeaveReq) (*emptypb.Empty, error) {
	e := m.Impl.SetGroupLeave(in.GroupId, in.IsDismiss)
	return &emptypb.Empty{}, e
}

//设置群组专属头衔
func (m *OnebotCliServerStub) SetGroupSpecialTitle(ctx context.Context, in *SetGroupSpecialTitleReq) (*emptypb.Empty, error) {
	e := m.Impl.SetGroupSpecialTitle(in.GroupId, in.UserId, in.Duration, in.SpecialTitle)
	return &emptypb.Empty{}, e
}

//处理加群请求／邀请
func (m *OnebotCliServerStub) SetGroupAddRequest(ctx context.Context, in *SetGroupAddRequestReq) (*emptypb.Empty, error) {
	e := m.Impl.SetGroupAddRequest(in.Flag, in.SubType, in.Reason, in.Approve)
	return &emptypb.Empty{}, e
}

//获取群信息
func (m *OnebotCliServerStub) GetGroupInfo(ctx context.Context, in *GetGrooupInfoReq) (*model.GroupGRPC, error) {
	r, e := m.Impl.GetGroupInfo(in.GroupId, in.NotCache)
	if e != nil {
		return nil, e
	}
	return r.ToGRPC(), e
}

//获取群列表
func (m *OnebotCliServerStub) GetGroupList(ctx context.Context, in *emptypb.Empty) (*GetGroupListResp, error) {
	r, e := m.Impl.GetGroupList()
	if e != nil {
		return nil, e
	}
	return &GetGroupListResp{Data: model.GroupArray2GroupGRPCArray(r)}, e
}

//获取群成员信息
func (m *OnebotCliServerStub) GetGroupMemberInfo(ctx context.Context, in *GetGroupMemberInfoReq) (*model.GroupMemberGRPC, error) {
	r, e := m.Impl.GetGroupMemberInfo(in.GroupId, in.UserId, in.NoCache)
	if e != nil {
		return nil, e
	}
	return r.ToGRPC(), nil
}

//获取群成员列表
func (m *OnebotCliServerStub) GetGroupMemberListInfo(ctx context.Context, in *emptypb.Empty) (*GetGroupMemberListInfoResp, error) {
	r, e := m.Impl.GetGroupMemberListInfo()
	if e != nil {
		return nil, e
	}
	return &GetGroupMemberListInfoResp{Data: model.GroupMemberArray2GroupMemberGRPCArray(r)}, nil
}

//获取群荣誉信息
func (m *OnebotCliServerStub) GetGroupHonorInfo(ctx context.Context, in *GetGroupHonorInfoReq) (*model.GroupHonorInfoGRPC, error) {
	r, e := m.Impl.GetGroupHonorInfo(in.GroupId, in.HonorType)
	if e != nil {
		return nil, e
	}
	return r.ToGRPC(), nil
}
func (m *OnebotCliServerStub) CleanCache(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, m.Impl.CleanCache()
}
func (m *OnebotCliServerStub) SetRestart(ctx context.Context, in *wrapperspb.Int64Value) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, m.Impl.SetRestart(in.Value)
}
func (m *OnebotCliServerStub) GetVersionInfo(ctx context.Context, in *emptypb.Empty) (*model.VersionInfoDataGRPC, error) {
	r, e := m.Impl.GetVersionInfo()
	if e != nil {
		return nil, e
	}
	return r.ToGRPC(), nil
}
func (m *OnebotCliServerStub) GetStatus(ctx context.Context, in *emptypb.Empty) (*model.StatusInfoDataGRPC, error) {
	r, e := m.Impl.GetStatus()
	if e != nil {
		return nil, e
	}
	return r.ToGRPC(), nil
}
func (m *OnebotCliServerStub) CanSendImage(ctx context.Context, in *emptypb.Empty) (*wrapperspb.BoolValue, error) {
	r, e := m.Impl.CanSendImage()
	if e != nil {
		return nil, e
	}
	return &wrapperspb.BoolValue{Value: r}, nil
}
func (m *OnebotCliServerStub) CanSendRecord(ctx context.Context, in *emptypb.Empty) (*wrapperspb.BoolValue, error) {
	r, e := m.Impl.CanSendRecord()
	if e != nil {
		return nil, e
	}
	return &wrapperspb.BoolValue{Value: r}, nil
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

func (m *OnebotCliClientStub) SendLike(userId int64, times int64) error {
	// 转发
	_, err := m.Client.SendLike(context.Background(), &SendLikeReq{UserId: userId, Times: times})
	return err
}

//处理加好友请求
func (m *OnebotCliClientStub) SetFriendAddRequest(flag string, approve bool, remark string) error {
	_, err := m.Client.SetFriendAddRequest(
		context.Background(),
		&SetFriendAddRequestReq{Flag: flag, Approve: approve, Remark: remark},
	)
	return err
}

//获取好友列表
func (m *OnebotCliClientStub) GetFriendList() ([]*model.Account, error) {
	r, err := m.Client.GetFriendList(
		context.Background(),
		&emptypb.Empty{},
	)
	if err != nil {
		return nil, err
	}
	return model.AccountGRPCArray2AccountArray(r.Data), err
}

// 群组踢人
func (m *OnebotCliClientStub) SetGroupKick(groupId, userId int64, rejectAddRequest bool) error {
	_, err := m.Client.SetGroupKick(
		context.Background(),
		&SetGroupKickReq{GroupId: groupId, UserId: userId, RejectRddRequest: rejectAddRequest},
	)
	return err
}

// 群组禁言
func (m *OnebotCliClientStub) SetGroupBan(groupId, userId, duration int64) error {
	_, err := m.Client.SetGroupBan(
		context.Background(),
		&SetGroupBanReq{GroupId: groupId, UserId: userId, Duration: duration},
	)
	return err
}

// 群组匿名用户禁言
func (m *OnebotCliClientStub) SetGroupAnonymousBan(groupId, duration int64, anonymousFlag string) error {
	_, err := m.Client.SetGroupAnonymousBan(
		context.Background(),
		&SetGroupAnonymousBanReq{GroupId: groupId, AnonymousFlag: anonymousFlag, Duration: duration},
	)
	return err
}

//群组全员禁言
func (m *OnebotCliClientStub) SetGroupWholeBan(groupId int64, enable bool) error {
	_, err := m.Client.SetGroupWholeBan(
		context.Background(),
		&SetGroupWholeBanReq{GroupId: groupId, Enable: enable},
	)
	return err
}

//群组设置管理员
func (m *OnebotCliClientStub) SetGroupAdmin(groupId, userId int64, enable bool) error {
	_, err := m.Client.SetGroupAdmin(
		context.Background(),
		&SetGroupAdminReq{GroupId: groupId, Enable: enable, UserId: userId},
	)
	return err
}

//群组匿名
func (m *OnebotCliClientStub) SetGroupAnonymous(groupId int64, enable bool) error {
	_, err := m.Client.SetGroupAnonymous(
		context.Background(),
		&SetGroupAnonymousReq{GroupId: groupId, Enable: enable},
	)
	return err
}

//设置群名片（群备注）
func (m *OnebotCliClientStub) SetGroupCard(groupId, userId int64, card string) error {
	_, err := m.Client.SetGroupCard(
		context.Background(),
		&SetGroupCardReq{GroupId: groupId, UserId: userId, Card: card},
	)
	return err
}

//设置群名
func (m *OnebotCliClientStub) SetGroupName(groupId int64, groupName string) error {
	_, err := m.Client.SetGroupName(
		context.Background(),
		&SetGroupNameReq{GroupId: groupId, GroupName: groupName},
	)
	return err
}

//退出群组
func (m *OnebotCliClientStub) SetGroupLeave(groupId int64, isDismiss bool) error {
	_, err := m.Client.SetGroupLeave(
		context.Background(),
		&SetGroupLeaveReq{GroupId: groupId, IsDismiss: isDismiss},
	)
	return err
}

//设置群组专属头衔
func (m *OnebotCliClientStub) SetGroupSpecialTitle(groupId, userId, duration int64, specialTitle string) error {
	_, err := m.Client.SetGroupSpecialTitle(
		context.Background(),
		&SetGroupSpecialTitleReq{GroupId: groupId, UserId: userId, SpecialTitle: specialTitle},
	)
	return err
}

//处理加群请求／邀请
func (m *OnebotCliClientStub) SetGroupAddRequest(flag, subType, reason string, approve bool) error {
	_, err := m.Client.SetGroupAddRequest(
		context.Background(),
		&SetGroupAddRequestReq{Flag: flag, SubType: subType, Reason: reason},
	)
	return err
}

//获取群信息
func (m *OnebotCliClientStub) GetGroupInfo(groupId int64, noCache bool) (*model.Group, error) {
	r, err := m.Client.GetGroupInfo(
		context.Background(),
		&GetGrooupInfoReq{GroupId: groupId, NotCache: noCache},
	)
	if err != nil {
		return nil, err
	}
	return r.ToStruct(), err
}

//获取群列表
func (m *OnebotCliClientStub) GetGroupList() ([]*model.Group, error) {
	r, err := m.Client.GetGroupList(
		context.Background(),
		&emptypb.Empty{},
	)
	if err != nil {
		return nil, err
	}
	return model.GroupGRPCArray2GroupArray(r.Data), err
}

//获取群成员信息
func (m *OnebotCliClientStub) GetGroupMemberInfo(groupId, userId int64, noCache bool) (*model.GroupMember, error) {
	r, err := m.Client.GetGroupMemberInfo(
		context.Background(),
		&GetGroupMemberInfoReq{
			GroupId: groupId,
			UserId:  userId,
			NoCache: noCache,
		},
	)
	if err != nil {
		return nil, err
	}
	return r.ToStruct(), err
}

//获取群成员列表
func (m *OnebotCliClientStub) GetGroupMemberListInfo() ([]*model.GroupMember, error) {
	r, err := m.Client.GetGroupMemberListInfo(
		context.Background(),
		&emptypb.Empty{},
	)
	if err != nil {
		return nil, err
	}
	return model.GroupMemberGRPCArray2GroupMemberArray(r.Data), err
}

//获取群荣誉信息
func (m *OnebotCliClientStub) GetGroupHonorInfo(groupId int64, honorType string) (*model.GroupHonorInfo, error) {
	r, err := m.Client.GetGroupHonorInfo(
		context.Background(),
		&GetGroupHonorInfoReq{
			GroupId:   groupId,
			HonorType: honorType,
		},
	)
	if err != nil {
		return nil, err
	}
	return r.ToStruct(), err
}

func (m *OnebotCliClientStub) CleanCache() error {
	_, err := m.Client.CleanCache(
		context.Background(),
		&emptypb.Empty{},
	)
	return err
}

func (m *OnebotCliClientStub) SetRestart(delay int64) error {
	_, err := m.Client.SetRestart(
		context.Background(),
		&wrapperspb.Int64Value{Value: delay},
	)
	return err
}

func (m *OnebotCliClientStub) GetVersionInfo() (*model.VersionInfoData, error) {
	r, err := m.Client.GetVersionInfo(
		context.Background(),
		&emptypb.Empty{},
	)
	if err != nil {
		return nil, err
	}
	return r.ToStruct(), err
}

func (m *OnebotCliClientStub) GetStatus() (*model.StatusInfoData, error) {
	r, err := m.Client.GetStatus(
		context.Background(),
		&emptypb.Empty{},
	)
	if err != nil {
		return nil, err
	}
	return r.ToStruct(), err
}

func (m *OnebotCliClientStub) CanSendImage() (bool, error) {
	r, err := m.Client.CanSendImage(
		context.Background(),
		&emptypb.Empty{},
	)
	if err != nil {
		return false, err
	}
	return r.Value, err
}

func (m *OnebotCliClientStub) CanSendRecord() (bool, error) {
	r, err := m.Client.CanSendRecord(
		context.Background(),
		&emptypb.Empty{},
	)
	if err != nil {
		return false, err
	}
	return r.Value, err
}

//插件实现GRPC的接口
func (p *OnebotCliGRPCPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &OnebotCliClientStub{Client: NewOnebotGrpcCliClient(c)}, nil
}
