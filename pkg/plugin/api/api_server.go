package api

import (
	"context"

	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/api/account"
	"github.com/dezhishen/onebot-sdk/pkg/api/cqhttp"
	"github.com/dezhishen/onebot-sdk/pkg/api/file"
	"github.com/dezhishen/onebot-sdk/pkg/api/friend"
	"github.com/dezhishen/onebot-sdk/pkg/api/friendopt"
	"github.com/dezhishen/onebot-sdk/pkg/api/group"
	"github.com/dezhishen/onebot-sdk/pkg/api/groupopt"
	"github.com/dezhishen/onebot-sdk/pkg/api/image"
	"github.com/dezhishen/onebot-sdk/pkg/api/message"
	"github.com/dezhishen/onebot-sdk/pkg/api/record"
	"github.com/dezhishen/onebot-sdk/pkg/api/request"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type OnebotAPiClientInterfaceGRPCServiceServer interface {
	account.OnebotApiAccountGRPCServiceServer
	cqhttp.OnebotApiCqhttpGRPCServiceServer
	file.OnebotApiFileGRPCServiceServer
	friend.OnebotApiFriendGRPCServiceServer
	friendopt.OnebotApiFriendOptGRPCServiceServer
	group.OnebotApiGroupGRPCServiceServer
	groupopt.OnebotApiGroupOptGRPCServiceServer
	image.OnebotApiImageGRPCServiceServer
	message.OnebotApiMessageGRPCServiceServer
	record.OnebotApiRecordGRPCServiceServer
	request.OnebotApiRequestGRPCServiceServer
}

func RegisterOnebotApiGRPCServiceServer(s *grpc.Server, srv OnebotAPiClientInterfaceGRPCServiceServer) {
	// api.RegisterOnebotGrpcCliServer(s, srv)
	account.RegisterOnebotApiAccountGRPCServiceServer(s, srv)
	cqhttp.RegisterOnebotApiCqhttpGRPCServiceServer(s, srv)
	file.RegisterOnebotApiFileGRPCServiceServer(s, srv)
	friend.RegisterOnebotApiFriendGRPCServiceServer(s, srv)
	friendopt.RegisterOnebotApiFriendOptGRPCServiceServer(s, srv)
	group.RegisterOnebotApiGroupGRPCServiceServer(s, srv)
	groupopt.RegisterOnebotApiGroupOptGRPCServiceServer(s, srv)
	image.RegisterOnebotApiImageGRPCServiceServer(s, srv)
	message.RegisterOnebotApiMessageGRPCServiceServer(s, srv)
	record.RegisterOnebotApiRecordGRPCServiceServer(s, srv)
	request.RegisterOnebotApiRequestGRPCServiceServer(s, srv)
}

type OnebotApiClient struct{}

type OnebotApiServerStub struct {
	plugin.Plugin
	// 具体实现，仅当业务接口实现基于Go时该字段有用
	Impl api.OnebotApiClientInterface
}

func (svc *OnebotApiServerStub) GetLoginInfo(ctx context.Context, in *emptypb.Empty) (*model.AccountResultGRPC, error) {
	res, err := svc.Impl.GetLoginInfo()
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

func (svc *OnebotApiServerStub) SetQQProfile(ctx context.Context, in *model.QQProfileGRPC) (*emptypb.Empty, error) {
	if err := svc.Impl.SetQQProfile(in.ToStruct()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (svc *OnebotApiServerStub) GetModelShow(ctx context.Context, in *emptypb.Empty) (*model.ModelShowResultGRPC, error) {
	res, err := svc.Impl.GetModelShow()
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

func (svc *OnebotApiServerStub) SetModelShow(ctx context.Context, in *account.SetModelShowRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetModelShow(in.GetModel(), in.GetModelshow()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (svc *OnebotApiServerStub) GetOnlineClients(ctx context.Context, in *emptypb.Empty) (*model.OnlineClientsResultGRPC, error) {
	res, err := svc.Impl.GetOnlineClients()
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// 获取 Cookies
// get_cookies
// domain 指定域名
func (svc *OnebotApiServerStub) GetCookies(ctx context.Context, in *wrapperspb.StringValue) (*model.CookiesResultGRPC, error) {
	res, err := svc.Impl.GetCookies(in.GetValue())
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// get_csrf_token
func (svc *OnebotApiServerStub) GetCsrfToken(ctx context.Context, in *emptypb.Empty) (*model.CsrfTokenResultGRPC, error) {
	res, err := svc.Impl.GetCsrfToken()
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// 获取 QQ 相关接口凭证
// get_credentials
func (svc *OnebotApiServerStub) GetCredentials(ctx context.Context, in *emptypb.Empty) (*model.CredentialsResultGRPC, error) {
	res, err := svc.Impl.GetCredentials()
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// 获取版本信息
// get_version_info
func (svc *OnebotApiServerStub) GetVersionInfo(ctx context.Context, in *emptypb.Empty) (*model.VersionInfoResultGRPC, error) {
	res, err := svc.Impl.GetVersionInfo()
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// 获取状态
// get_status
func (svc *OnebotApiServerStub) GetStatus(ctx context.Context, in *emptypb.Empty) (*model.StatusResultGRPC, error) {
	res, err := svc.Impl.GetStatus()
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// 重启 Go-CqHttp
// set_restart
// delay 要延迟的毫秒数, 如果默认情况下无法重启, 可以尝试设置延迟为 2000 左右
func (svc *OnebotApiServerStub) SetRestart(ctx context.Context, in *wrapperspb.Int32Value) (*emptypb.Empty, error) {
	if err := svc.Impl.SetRestart(in.GetValue()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 清理缓存
// clean_cache
func (svc *OnebotApiServerStub) CleanCache(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	if err := svc.Impl.CleanCache(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 重载事件过滤器
// reload_event_filter
// file 事件过滤器文件路径
func (svc *OnebotApiServerStub) ReloadEventFilter(ctx context.Context, in *wrapperspb.StringValue) (*emptypb.Empty, error) {
	if err := svc.Impl.ReloadEventFilter(in.GetValue()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 下载文件到缓存目录
// download_file
// url 文件链接
// thread_count 线程数
// headers 请求头
func (svc *OnebotApiServerStub) DownloadFile(ctx context.Context, in *cqhttp.DownloadFileRequestGRPC) (*model.DownloadFileResultGRPC, error) {
	res, err := svc.Impl.DownloadFile(in.GetUrl(), in.GetThreadCount(), in.GetHeaders())
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// 检查链接安全性
// check_url_safely
// url 链接
func (svc *OnebotApiServerStub) CheckUrlSafely(ctx context.Context, in *wrapperspb.StringValue) (*model.CheckUrlSafelyResultGRPC, error) {
	res, err := svc.Impl.CheckUrlSafely(in.GetValue())
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// 获取中文分词 ( 隐藏 API )
// .get_word_slices
// content 内容
func (svc *OnebotApiServerStub) HiddenGetWordSlices(ctx context.Context, in *wrapperspb.StringValue) (*model.WordSlicesResultGRPC, error) {
	res, err := svc.Impl.HiddenGetWordSlices(in.GetValue())
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// 对事件执行快速操作 ( 隐藏 API )
// .handle_quick_operation
// context 事件上下文
// operation 操作
func (svc *OnebotApiServerStub) HiddenHandleQuickOperation(ctx context.Context, in *cqhttp.HiddenHandleQuickOperationRequestGRPC) (*emptypb.Empty, error) {
	if err := svc.Impl.HiddenHandleQuickOperation(in.GetContext(), in.GetOperation()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 上传群文件
// upload_group_file
// groupId 群号
// file 文件路径
// name 文件名
// folder 群文件目录ID
func (svc *OnebotApiServerStub) UploadGroupFile(ctx context.Context, in *file.UploadGroupFileRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.UploadGroupFile(in.GetGroupId(), in.GetFile(), in.GetName(), in.GetFolder()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 删除群文件
// delete_group_file
// groupId 群号
// file_id 文件ID 参考 File 对象
// busid 文件类型 参考 File 对象
// DeleteGroupFile(groupId int64, fileId string, busid int32) error
func (svc *OnebotApiServerStub) DeleteGroupFile(ctx context.Context, in *file.DeleteGroupFileRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.DeleteGroupFile(in.GetGroupId(), in.GetFileId(), in.GetBusid()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 创建群文件目录
// create_group_file_folder
// groupId 群号
// name 目录名
// parentId 父目录ID
// CreateGroupFileFolder(groupId int64, name string, parentId string) error
func (svc *OnebotApiServerStub) CreateGroupFileFolder(ctx context.Context, in *file.CreateGroupFileFolderRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.CreateGroupFileFolder(in.GetGroupId(), in.GetName(), in.GetParentId()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 删除群文件目录
// delete_group_folder
// groupId 群号
// folder_id 文件夹ID 参考 Folder 对象
// DeleteGroupFolder(groupId int64, folderId string) error
func (svc *OnebotApiServerStub) DeleteGroupFolder(ctx context.Context, in *file.DeleteGroupFolderRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.DeleteGroupFolder(in.GetGroupId(), in.GetFolderId()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 获取群文件系统信息
// get_group_file_system_info
// groupId 群号
// GetGroupFileSystemInfo(groupId int64) (*model.GroupFileSystemInfoResult, error)
func (svc *OnebotApiServerStub) GetGroupFileSystemInfo(ctx context.Context, in *wrapperspb.Int64Value) (*model.GroupFileSystemInfoResultGRPC, error) {
	res, err := svc.Impl.GetGroupFileSystemInfo(in.GetValue())
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// 获取群根目录文件列表
// get_group_root_files
// groupId 群号
// GetGroupRootFiles(groupId int64) (*model.GroupFilesResult, error)
func (svc *OnebotApiServerStub) GetGroupRootFiles(ctx context.Context, in *wrapperspb.Int64Value) (*model.GroupFilesResultGRPC, error) {
	res, err := svc.Impl.GetGroupRootFiles(in.GetValue())
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// get_group_files_by_folder
// 获取群文件列表
// groupId 群号
// folder_id 文件夹ID
// GetGroupFilesByFolder(groupId int64, folderId string) (*model.GroupFilesResult, error)
func (svc *OnebotApiServerStub) GetGroupFilesByFolder(ctx context.Context, in *file.GroupFilesByFolderRequest) (*model.GroupFilesResultGRPC, error) {
	res, err := svc.Impl.GetGroupFilesByFolder(in.GetGroupId(), in.GetFolderId())
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// 获取群文件资源链接
// get_group_file_url
// groupId 群号
// file_id 文件ID 参考 File 对象
// busid 文件类型 参考 File 对象
// GetGroupFileUrl(groupId int64, fileId string, busid int32) (*model.FileUrlResult, error)
func (svc *OnebotApiServerStub) GetGroupFileUrl(ctx context.Context, in *file.GetGroupFileUrlRequest) (*model.FileUrlResultGRPC, error) {
	res, err := svc.Impl.GetGroupFileUrl(in.GetGroupId(), in.GetFileId(), in.GetBusid())
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// 上传好友文件
// upload_private_file
// userId 用户ID
// file 文件路径
// name 文件名
// UploadPrivateFile(userId int64, file string, name string) error
func (svc *OnebotApiServerStub) UploadPrivateFile(ctx context.Context, in *file.UploadPrivateFileRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.UploadPrivateFile(in.GetUserId(), in.GetFile(), in.GetName()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 获取陌生人信息
// get_stranger_info
// user_id: 对方 QQ 号
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (svc *OnebotApiServerStub) GetStrangerInfo(ctx context.Context, in *friend.GetStrangerInfoRequest) (*model.StrangerInfoResultGRPC, error) {
	res, err := svc.Impl.GetStrangerInfo(in.GetUserId(), in.GetNoCache())
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.ToGRPC(), nil
}

// 获取好友列表
// get_friend_list
func (svc *OnebotApiServerStub) GetFriendList(ctx context.Context, in *emptypb.Empty) (*model.FriendListResultGRPC, error) {
	if res, err := svc.Impl.GetFriendList(); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 获取单向好友列表
// get_unidirectional_friend_list
func (svc *OnebotApiServerStub) GetUnidirectionalFriendList(ctx context.Context, in *emptypb.Empty) (*model.FriendListResultGRPC, error) {
	if res, err := svc.Impl.GetUnidirectionalFriendList(); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 删除好友
// delete_friend
// user_id: int64 对方 QQ 号
func (svc *OnebotApiServerStub) DeleteFriend(ctx context.Context, in *wrapperspb.Int64Value) (*emptypb.Empty, error) {
	if err := svc.Impl.DeleteFriend(in.GetValue()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 删除单向好友
// delete_unidirectional_friend
// user_id: int64 对方 QQ 号
func (svc *OnebotApiServerStub) DeleteUnidirectionalFriend(ctx context.Context, in *wrapperspb.Int64Value) (*emptypb.Empty, error) {
	if err := svc.Impl.DeleteUnidirectionalFriend(in.GetValue()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 获取群信息
// get_group_info
// group_id: 群号
// no_cache: 是否使用缓存（使用缓存可能更新不及时，但响应更快）
func (svc *OnebotApiServerStub) GetGroupInfo(ctx context.Context, in *group.GetGroupInfoRequest) (*model.GroupInfoResultGRPC, error) {
	if res, err := svc.Impl.GetGroupInfo(in.GetGroupId(), in.GetNoCache()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 获取群列表
// get_group_list
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (svc *OnebotApiServerStub) GetGroupList(ctx context.Context, in *wrapperspb.BoolValue) (*model.GroupListResultGRPC, error) {
	if res, err := svc.Impl.GetGroupList(in.GetValue()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 获取群成员信息
// get_group_member_info
// group_id: 群号
// user_id: QQ 号
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (svc *OnebotApiServerStub) GetGroupMemberInfo(ctx context.Context, in *group.GetGroupMemberInfoRequest) (*model.GroupMemberInfoResultGRPC, error) {
	if res, err := svc.Impl.GetGroupMemberInfo(in.GetGroupId(), in.GetUserId(), in.GetNoCache()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 获取群成员列表
// get_group_member_list
// group_id: 群号
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (svc *OnebotApiServerStub) GetGroupMemberList(ctx context.Context, in *group.GetGroupMemberListRequest) (*model.GroupMemberListResultGRPC, error) {
	if res, err := svc.Impl.GetGroupMemberList(in.GetGroupId(), in.GetNoCache()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 获取群荣誉信息
// get_group_honor_info
// group_id: 群号
// type: 群荣誉类型，目前支持 talkative（群聊之火）、performer（群聊炽焰）、legend（群聊传说）、strong_newbie（群聊新星）、emotion（群表情之火）
func (svc *OnebotApiServerStub) GetGroupHonorInfo(ctx context.Context, in *group.GetGroupHonorInfoRequest) (*model.GroupHonorInfoResultGRPC, error) {
	if res, err := svc.Impl.GetGroupHonorInfo(in.GetGroupId(), in.GetType()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 获取群系统消息
// get_group_system_msg
func (svc *OnebotApiServerStub) GetGroupSystemMsg(ctx context.Context, in *emptypb.Empty) (*model.GroupSystemMsgResultGRPC, error) {
	if res, err := svc.Impl.GetGroupSystemMsg(); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 获取精华消息列表
// get_essence_msg_list
// group_id: 群号
func (svc *OnebotApiServerStub) GetEssenceMsgList(ctx context.Context, in *wrapperspb.Int64Value) (*model.EssenceMsgListResultGRPC, error) {
	if res, err := svc.Impl.GetEssenceMsgList(in.GetValue()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 获取群 @全体成员 剩余次数
// get_group_at_all_remain
// group_id: 群号
func (svc *OnebotApiServerStub) GetGroupAtAllRemain(ctx context.Context, in *wrapperspb.Int64Value) (*model.GroupAtAllRemainResultGRPC, error) {
	if res, err := svc.Impl.GetGroupAtAllRemain(in.GetValue()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 设置群名
// set_group_name
// groupId 群号
// groupName 群名称
func (svc *OnebotApiServerStub) SetGroupName(ctx context.Context, in *groupopt.SetGroupNameRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetGroupName(in.GetGroupId(), in.GetGroupName()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 设置群头像
// set_group_portrait
// groupId 群号
// file 图片文件路径
// cache 缓存时间
func (svc *OnebotApiServerStub) SetGroupPortrait(ctx context.Context, in *groupopt.SetGroupPortraitRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetGroupPortrait(in.GetGroupId(), in.GetFile(), int(in.GetCache())); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// set_group_admin
// 设置群管理员
// groupId 群号
// userId QQ 号
// enable true 为设置，false 为取消
func (svc *OnebotApiServerStub) SetGroupAdmin(ctx context.Context, in *groupopt.SetGroupAdminRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetGroupAdmin(in.GetGroupId(), in.GetUserId(), in.GetEnable()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 设置群名片
// set_group_card
// groupId 群号
// userId QQ 号
// card 群名片内容，不填或空字符串表示删除群名片
func (svc *OnebotApiServerStub) SetGroupCard(ctx context.Context, in *groupopt.SetGroupCardRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetGroupCard(in.GetGroupId(), in.GetUserId(), in.GetCard()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 设置群头衔
// set_group_special_title
// groupId 群号
// userId QQ 号
// specialTitle 头衔，不填或空字符串表示删除群头衔
// duration 专属头衔有效期, 单位秒, -1 表示永久, 不过此项似乎没有效果, 可能是只有某些特殊的时间长度有效, 有待测试
func (svc *OnebotApiServerStub) SetGroupSpecialTitle(ctx context.Context, in *groupopt.SetGroupSpecialTitleRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetGroupSpecialTitle(in.GetGroupId(), in.GetUserId(), in.GetSpecialTitle(), in.GetDuration()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 禁言群成员
// set_group_ban
// groupId 群号
// userId QQ 号
// duration 禁言时长，单位秒，0 表示取消禁言
func (svc *OnebotApiServerStub) SetGroupBan(ctx context.Context, in *groupopt.SetGroupBanRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetGroupBan(in.GetGroupId(), in.GetUserId(), in.GetDuration()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 设置全群禁言
// set_group_whole_ban
// groupId 群号
// enable 是否禁言
func (svc *OnebotApiServerStub) SetGroupWholeBan(ctx context.Context, in *groupopt.SetGroupWholeBanequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetGroupWholeBan(in.GetGroupId(), in.GetEnable()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 禁言群匿名成员
// set_group_anonymous_ban
// groupId 群号
// anonymous 可选, 要禁言的匿名用户对象（群消息上报的 anonymous 字段）
// anonymousFlag 可选, 要禁言的匿名用户的 flag（需从群消息上报的数据中获得）
// 上面的 anonymous 和 anonymous_flag 两者任选其一传入即可, 若都传入, 则使用 anonymous。
// duration 禁言时长，单位秒，不能超过 30 天
func (svc *OnebotApiServerStub) SetGroupAnonymousBan(ctx context.Context, in *groupopt.SetGroupAnonymousBanRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetGroupAnonymousBan(in.GetGroupId(), in.GetAnonymous(), in.GetAnonymousFlag(), in.GetDuration()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 设置精华消息
// set_essence_msg
// message_id 消息 ID
func (svc *OnebotApiServerStub) SetEssenceMsg(ctx context.Context, in *wrapperspb.Int64Value) (*emptypb.Empty, error) {
	if err := svc.Impl.SetEssenceMsg(in.GetValue()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 删除精华消息
// delete_essence_msg
// message_id 消息 ID
func (svc *OnebotApiServerStub) DeleteEssenceMsg(ctx context.Context, in *wrapperspb.Int64Value) (*emptypb.Empty, error) {
	if err := svc.Impl.DeleteEssenceMsg(in.GetValue()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 发送群签到
// send_group_sign
// groupId 群号
func (svc *OnebotApiServerStub) SendGroupSign(ctx context.Context, in *wrapperspb.Int64Value) (*emptypb.Empty, error) {
	if err := svc.Impl.SendGroupSign(in.GetValue()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 设置群匿名
// set_group_anonymous
// groupId 群号
// enable 是否允许匿名聊天
func (svc *OnebotApiServerStub) SetGroupAnonymous(ctx context.Context, in *groupopt.SetGroupAnonymousRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetGroupAnonymous(in.GetGroupId(), in.GetEnable()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 发送群公告
// _send_group_notice
// groupId 群号
// content 公告内容
// image 图片文件路径（可选）
func (svc *OnebotApiServerStub) SendGroupNotice(ctx context.Context, in *groupopt.SendGroupNoticeRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SendGroupNotice(in.GetGroupId(), in.GetContent(), in.GetImage()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 获取群公告
// _get_group_notice
// groupId 群号
func (svc *OnebotApiServerStub) GetGroupNotice(ctx context.Context, in *wrapperspb.Int64Value) (*model.GroupNoticeResultGRPC, error) {
	if res, err := svc.Impl.GetGroupNotice(in.GetValue()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 群组踢人
// set_group_kick
// groupId 群号
// userId QQ 号
// rejectAddRequest 是否拒绝此人的加群请求
func (svc *OnebotApiServerStub) SetGroupKick(ctx context.Context, in *groupopt.SetGroupKickRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetGroupKick(in.GetGroupId(), in.GetUserId(), in.GetRejectAddRequest()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 退出群组
// set_group_leave
// groupId 群号
// isDismiss 是否解散，如果登录号是群主，则仅在此项为 true 时能够解散
func (svc *OnebotApiServerStub) SetGroupLeave(ctx context.Context, in *groupopt.SetGroupLeaveRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetGroupLeave(in.GetGroupId(), in.GetIsDismiss()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 获取图片
// get_image
func (svc *OnebotApiServerStub) GetImage(ctx context.Context, in *wrapperspb.StringValue) (*model.ImageResultGRPC, error) {
	if res, err := svc.Impl.GetImage(in.GetValue()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 检查是否可以发送图片
// can_send_image
func (svc *OnebotApiServerStub) CanSendImage(ctx context.Context, in *emptypb.Empty) (*model.BoolYesOfResultGRPC, error) {
	if res, err := svc.Impl.CanSendImage(); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 图片ORC识别
// ocr_image
func (svc *OnebotApiServerStub) OcrImage(ctx context.Context, in *wrapperspb.StringValue) (*model.OcrImageResultGRPC, error) {
	if res, err := svc.Impl.OcrImage(in.GetValue()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 发送私信
// send_private_msg
// msg 消息
func (svc *OnebotApiServerStub) SendPrivateMsg(ctx context.Context, in *model.PrivateMsgGRPC) (*model.SendMessageResultGRPC, error) {
	if res, err := svc.Impl.SendPrivateMsg(in.ToStruct()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 发送群消息
// send_group_msg
// msg 消息
func (svc *OnebotApiServerStub) SendGroupMsg(ctx context.Context, in *model.GroupMsgGRPC) (*model.SendMessageResultGRPC, error) {
	if res, err := svc.Impl.SendGroupMsg(in.ToStruct()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 发送消息
func (svc *OnebotApiServerStub) SendMsg(ctx context.Context, in *model.MsgForSendGRPC) (*model.SendMessageResultGRPC, error) {
	if res, err := svc.Impl.SendMsg(in.ToStruct()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 获取消息
// get_msg
func (svc *OnebotApiServerStub) GetMsg(ctx context.Context, in *wrapperspb.Int64Value) (*model.MessageDataResultGRPC, error) {
	if res, err := svc.Impl.GetMsg(in.GetValue()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 撤回消息
// delete_msg
func (svc *OnebotApiServerStub) DeleteMsg(ctx context.Context, in *wrapperspb.Int64Value) (*emptypb.Empty, error) {
	if err := svc.Impl.DeleteMsg(in.GetValue()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 标记消息已读
// mark_msg_as_read
func (svc *OnebotApiServerStub) MarkMsgAsRead(ctx context.Context, in *wrapperspb.Int64Value) (*emptypb.Empty, error) {
	if err := svc.Impl.MarkMsgAsRead(in.GetValue()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 获取合并转发消息
// get_forward_msg
func (svc *OnebotApiServerStub) GetForwardMsg(ctx context.Context, in *wrapperspb.Int64Value) (*model.ForwardMessageDataResultGRPC, error) {
	if res, err := svc.Impl.GetForwardMsg(in.GetValue()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 发送合并转发 ( 群聊 )
// send_group_forward_msg
func (svc *OnebotApiServerStub) SendGroupForwardMsg(ctx context.Context, in *message.SendGroupForwardMsgRequest) (*model.SendForwardMessageDataResultGRPC, error) {
	if res, err := svc.Impl.SendGroupForwardMsg(in.GetGroupId(), model.MessageSegmentGRPCArray2MessageSegmentArray(in.GetMessages())); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 发送合并转发 ( 好友 )
// send_private_forward_msg
func (svc *OnebotApiServerStub) SendPrivateForwardMsg(ctx context.Context, in *message.SendPrivateForwardMsgGRPC) (*model.SendForwardMessageDataResultGRPC, error) {
	if res, err := svc.Impl.SendPrivateForwardMsg(in.GetUserId(), model.MessageSegmentGRPCArray2MessageSegmentArray(in.GetMessages())); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 获取群消息历史记录
// get_group_msg_history
func (svc *OnebotApiServerStub) GetGroupMsgHistory(ctx context.Context, in *message.GetGroupMsgHistoryRequest) (*model.MessagesResultGRPC, error) {
	if res, err := svc.Impl.GetGroupMsgHistory(in.GetMessageSeq(), in.GetGroupId()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 获取语音
// get_record
func (svc *OnebotApiServerStub) GetRecord(ctx context.Context, in *record.GetRecordRequest) (*model.RecordResultGRPC, error) {
	if res, err := svc.Impl.GetRecord(in.GetFile(), in.GetOutFormat()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 检查是否可以发送语音
// can_send_record
func (svc *OnebotApiServerStub) CanSendRecord(ctx context.Context, in *emptypb.Empty) (*model.BoolYesOfResultGRPC, error) {
	if res, err := svc.Impl.CanSendRecord(); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToGRPC(), nil
	}
	return nil, nil
}

// 处理加好友请求
// set_friend_add_request
// flag: 加好友请求的 flag（需从上报的数据中获得）
// approve: 是否同意请求
// remark: 添加后的好友备注（仅在同意时有效）
func (svc *OnebotApiServerStub) SetFriendAddRequest(ctx context.Context, in *request.SetFriendAddRequestRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetFriendAddRequest(in.GetFlag(), in.GetApprove(), in.GetRemark()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// 处理加群请求／邀请
// set_group_add_request
// flag: 加群请求的 flag（需从上报的数据中获得）
// sub_type: add 或 invite，请求类型（需要和上报消息中的 sub_type 字段相符）
// approve: 是否同意请求／邀请
// reason: 拒绝理由（仅在拒绝时有效）
func (svc *OnebotApiServerStub) SetGroupAddRequest(ctx context.Context, in *request.SetGroupAddRequestRequest) (*emptypb.Empty, error) {
	if err := svc.Impl.SetGroupAddRequest(in.GetFlag(), in.GetSubType(), in.GetApprove(), in.GetReason()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
