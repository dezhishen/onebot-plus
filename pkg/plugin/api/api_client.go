package api

import (
	"context"
	"errors"

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
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type OnebotAPiClientInterfaceGRPCServiceClient interface {
	account.OnebotApiAccountGRPCServiceClient
	cqhttp.OnebotApiCqhttpGRPCServiceClient
	file.OnebotApiFileGRPCServiceClient
	friend.OnebotApiFriendGRPCServiceClient
	friendopt.OnebotApiFriendOptGRPCServiceClient
	group.OnebotApiGroupGRPCServiceClient
	groupopt.OnebotApiGroupOptGRPCServiceClient
	image.OnebotApiImageGRPCServiceClient
	message.OnebotApiMessageGRPCServiceClient
	record.OnebotApiRecordGRPCServiceClient
	request.OnebotApiRequestGRPCServiceClient
}

type onebotAPiClientInterfaceGRPCServiceClient struct {
	account.OnebotApiAccountGRPCServiceClient
	cqhttp.OnebotApiCqhttpGRPCServiceClient
	file.OnebotApiFileGRPCServiceClient
	friend.OnebotApiFriendGRPCServiceClient
	friendopt.OnebotApiFriendOptGRPCServiceClient
	group.OnebotApiGroupGRPCServiceClient
	groupopt.OnebotApiGroupOptGRPCServiceClient
	image.OnebotApiImageGRPCServiceClient
	message.OnebotApiMessageGRPCServiceClient
	record.OnebotApiRecordGRPCServiceClient
	request.OnebotApiRequestGRPCServiceClient
}

func NewOnebotAPiClientInterfaceGRPCServiceClient(c *grpc.ClientConn) OnebotAPiClientInterfaceGRPCServiceClient {
	return &onebotAPiClientInterfaceGRPCServiceClient{
		OnebotApiAccountGRPCServiceClient:   account.NewOnebotApiAccountGRPCServiceClient(c),
		OnebotApiCqhttpGRPCServiceClient:    cqhttp.NewOnebotApiCqhttpGRPCServiceClient(c),
		OnebotApiFileGRPCServiceClient:      file.NewOnebotApiFileGRPCServiceClient(c),
		OnebotApiFriendGRPCServiceClient:    friend.NewOnebotApiFriendGRPCServiceClient(c),
		OnebotApiFriendOptGRPCServiceClient: friendopt.NewOnebotApiFriendOptGRPCServiceClient(c),
		OnebotApiGroupGRPCServiceClient:     group.NewOnebotApiGroupGRPCServiceClient(c),
		OnebotApiGroupOptGRPCServiceClient:  groupopt.NewOnebotApiGroupOptGRPCServiceClient(c),
		OnebotApiImageGRPCServiceClient:     image.NewOnebotApiImageGRPCServiceClient(c),
		OnebotApiMessageGRPCServiceClient:   message.NewOnebotApiMessageGRPCServiceClient(c),
		OnebotApiRecordGRPCServiceClient:    record.NewOnebotApiRecordGRPCServiceClient(c),
		OnebotApiRequestGRPCServiceClient:   request.NewOnebotApiRequestGRPCServiceClient(c),
	}
}

// 业务接口的实现，通过gRPC客户端转发请求给插件进程
type OnebotApiClientStub struct {
	Client OnebotAPiClientInterfaceGRPCServiceClient
}

// 获取登录号信息
// get_login_info
func (cli *OnebotApiClientStub) GetLoginInfo() (*model.AccountResult, error) {
	if res, err := cli.Client.GetLoginInfo(context.Background(), &emptypb.Empty{}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 设置登录号资料
// set_login_info
func (cli *OnebotApiClientStub) SetQQProfile(profile *model.QQProfile) error {
	if _, err := cli.Client.SetQQProfile(context.Background(), profile.ToGRPC()); err != nil {
		return err
	}
	return nil
}

// 获取在线机型
// _get_model_show
func (cli *OnebotApiClientStub) GetModelShow() (*model.ModelShowResult, error) {
	if res, err := cli.Client.GetModelShow(context.Background(), &emptypb.Empty{}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 设置在线机型
// _set_model_show
// model: 机型
// modelshow: 机型展示
func (cli *OnebotApiClientStub) SetModelShow(model string, modelshow string) error {
	req := &account.SetModelShowRequest{
		Model:     model,
		Modelshow: modelshow,
	}
	if _, err := cli.Client.SetModelShow(context.Background(), req); err != nil {
		return err
	}
	return nil
}

// 获取在线客户端列表
// get_online_clients
func (cli *OnebotApiClientStub) GetOnlineClients() (*model.OnlineClientsResult, error) {
	if res, err := cli.Client.GetOnlineClients(context.Background(), &emptypb.Empty{}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取 Cookies
// get_cookies
// domain 指定域名
func (cli *OnebotApiClientStub) GetCookies(domain string) (*model.CookiesResult, error) {
	if res, err := cli.Client.GetCookies(context.Background(), &wrapperspb.StringValue{
		Value: domain,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取 CSRF Token
// get_csrf_token
func (cli *OnebotApiClientStub) GetCsrfToken() (*model.CsrfTokenResult, error) {
	if res, err := cli.Client.GetCsrfToken(context.Background(), &emptypb.Empty{}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取 QQ 相关接口凭证
// get_credentials
func (cli *OnebotApiClientStub) GetCredentials() (*model.CredentialsResult, error) {
	if res, err := cli.Client.GetCredentials(context.Background(), &emptypb.Empty{}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取版本信息
// get_version_info
func (cli *OnebotApiClientStub) GetVersionInfo() (*model.VersionInfoResult, error) {
	if res, err := cli.Client.GetVersionInfo(context.Background(), &emptypb.Empty{}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取状态
// get_status
func (cli *OnebotApiClientStub) GetStatus() (*model.StatusResult, error) {
	if res, err := cli.Client.GetStatus(context.Background(), &emptypb.Empty{}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 重启 Go-CqHttp
// set_restart
// delay 要延迟的毫秒数, 如果默认情况下无法重启, 可以尝试设置延迟为 2000 左右
func (cli *OnebotApiClientStub) SetRestart(delay int32) error {
	if _, err := cli.Client.SetRestart(context.Background(), &wrapperspb.Int32Value{
		Value: delay,
	}); err != nil {
		return err
	}
	return nil
}

// 清理缓存
// clean_cache
func (cli *OnebotApiClientStub) CleanCache() error {
	if _, err := cli.Client.CleanCache(context.Background(), &emptypb.Empty{}); err != nil {
		return err
	}
	return nil
}

// 重载事件过滤器
// reload_event_filter
// file 事件过滤器文件路径
func (cli *OnebotApiClientStub) ReloadEventFilter(file string) error {
	if _, err := cli.Client.ReloadEventFilter(context.Background(), &wrapperspb.StringValue{
		Value: file,
	}); err != nil {
		return err
	}
	return nil
}

// 下载文件到缓存目录
// download_file
// url 文件链接
// thread_count 线程数
// headers 请求头
func (cli *OnebotApiClientStub) DownloadFile(url string, threadCount int32, headers []string) (*model.DownloadFileResult, error) {
	req := &cqhttp.DownloadFileRequestGRPC{
		Url:         url,
		ThreadCount: threadCount,
		Headers:     headers,
	}
	if res, err := cli.Client.DownloadFile(context.Background(), req); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 检查链接安全性
// check_url_safely
// url 链接
func (cli *OnebotApiClientStub) CheckUrlSafely(url string) (*model.CheckUrlSafelyResult, error) {
	if res, err := cli.Client.CheckUrlSafely(context.Background(), &wrapperspb.StringValue{
		Value: url,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取中文分词 ( 隐藏 API )
// .get_word_slices
// content 内容
func (cli *OnebotApiClientStub) HiddenGetWordSlices(content string) (*model.WordSlicesResult, error) {
	if res, err := cli.Client.HiddenGetWordSlices(context.Background(), &wrapperspb.StringValue{
		Value: content,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 对事件执行快速操作 ( 隐藏 API )
// .handle_quick_operation
// context 事件上下文
// operation 操作
func (cli *OnebotApiClientStub) HiddenHandleQuickOperation(ctx interface{}, operation interface{}) error {
	panic("not implemented") // TODO: Implement
}

// 上传群文件
// upload_group_file
// groupId 群号
// file 文件路径
// name 文件名
// folder 群文件目录ID
func (cli *OnebotApiClientStub) UploadGroupFile(groupId int64, _file string, name string, folder string) error {
	if _, err := cli.Client.UploadGroupFile(context.Background(), &file.UploadGroupFileRequest{
		GroupId: groupId,
		File:    _file,
		Name:    name,
		Folder:  folder,
	}); err != nil {
		return err
	}
	return nil
}

// 删除群文件
// delete_group_file
// groupId 群号
// file_id 文件ID 参考 File 对象
// busid 文件类型 参考 File 对象
func (cli *OnebotApiClientStub) DeleteGroupFile(groupId int64, fileId string, busid int32) error {
	if _, err := cli.Client.DeleteGroupFile(context.Background(), &file.DeleteGroupFileRequest{
		GroupId: groupId,
		FileId:  fileId,
		Busid:   busid,
	}); err != nil {
		return err
	}
	return nil
}

// 创建群文件目录
// create_group_file_folder
// groupId 群号
// name 目录名
// parentId 父目录ID
func (cli *OnebotApiClientStub) CreateGroupFileFolder(groupId int64, name string, parentId string) error {
	if _, err := cli.Client.CreateGroupFileFolder(context.Background(), &file.CreateGroupFileFolderRequest{
		GroupId:  groupId,
		Name:     name,
		ParentId: parentId,
	}); err != nil {
		return err
	}
	return nil
}

// 删除群文件目录
// delete_group_folder
// groupId 群号
// folder_id 文件夹ID 参考 Folder 对象
func (cli *OnebotApiClientStub) DeleteGroupFolder(groupId int64, folderId string) error {
	if _, err := cli.Client.DeleteGroupFolder(context.Background(), &file.DeleteGroupFolderRequest{
		GroupId:  groupId,
		FolderId: folderId,
	}); err != nil {
		return err
	}
	return nil
}

// 获取群文件系统信息
// get_group_file_system_info
// groupId 群号
func (cli *OnebotApiClientStub) GetGroupFileSystemInfo(groupId int64) (*model.GroupFileSystemInfoResult, error) {
	if res, err := cli.Client.GetGroupFileSystemInfo(context.Background(), &wrapperspb.Int64Value{
		Value: groupId,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取群根目录文件列表
// get_group_root_files
// groupId 群号
func (cli *OnebotApiClientStub) GetGroupRootFiles(groupId int64) (*model.GroupFilesResult, error) {
	if res, err := cli.Client.GetGroupRootFiles(context.Background(), &wrapperspb.Int64Value{
		Value: groupId,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// get_group_files_by_folder
// 获取群文件列表
// groupId 群号
// folder_id 文件夹ID
func (cli *OnebotApiClientStub) GetGroupFilesByFolder(groupId int64, folderId string) (*model.GroupFilesResult, error) {
	if res, err := cli.Client.GetGroupFilesByFolder(context.Background(), &file.GroupFilesByFolderRequest{
		GroupId:  groupId,
		FolderId: folderId,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取群文件资源链接
// get_group_file_url
// groupId 群号
// file_id 文件ID 参考 File 对象
// busid 文件类型 参考 File 对象
func (cli *OnebotApiClientStub) GetGroupFileUrl(groupId int64, fileId string, busid int32) (*model.FileUrlResult, error) {
	if res, err := cli.Client.GetGroupFileUrl(context.Background(), &file.GetGroupFileUrlRequest{
		GroupId: groupId,
		FileId:  fileId,
		Busid:   busid,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 上传好友文件
// upload_private_file
// userId 用户ID
// file 文件路径
// name 文件名
func (cli *OnebotApiClientStub) UploadPrivateFile(userId int64, _file string, name string) error {
	if _, err := cli.Client.UploadPrivateFile(context.Background(), &file.UploadPrivateFileRequest{
		UserId: userId,
		File:   _file,
		Name:   name,
	}); err != nil {
		return err
	}
	return nil
}

// 获取陌生人信息
// get_stranger_info
// user_id: 对方 QQ 号
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (cli *OnebotApiClientStub) GetStrangerInfo(userId int64, noCache bool) (*model.StrangerInfoResult, error) {
	if res, err := cli.Client.GetStrangerInfo(context.Background(), &friend.GetStrangerInfoRequest{
		UserId:  userId,
		NoCache: noCache,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取好友列表
// get_friend_list
func (cli *OnebotApiClientStub) GetFriendList() (*model.FriendListResult, error) {
	if res, err := cli.Client.GetFriendList(context.Background(), &emptypb.Empty{}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取单向好友列表
// get_unidirectional_friend_list
func (cli *OnebotApiClientStub) GetUnidirectionalFriendList() (*model.FriendListResult, error) {
	if res, err := cli.Client.GetUnidirectionalFriendList(context.Background(), &emptypb.Empty{}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 删除好友
// delete_friend
// user_id: 对方 QQ 号
func (cli *OnebotApiClientStub) DeleteFriend(userId int64) error {
	if _, err := cli.Client.DeleteFriend(context.Background(), &wrapperspb.Int64Value{
		Value: userId,
	}); err != nil {
		return err
	}
	return nil
}

// 删除单向好友
// delete_unidirectional_friend
// user_id: 对方 QQ 号
func (cli *OnebotApiClientStub) DeleteUnidirectionalFriend(userId int64) error {
	if _, err := cli.Client.DeleteUnidirectionalFriend(context.Background(), &wrapperspb.Int64Value{
		Value: userId,
	}); err != nil {
		return err
	}
	return nil
}

// 获取群信息
// get_group_info
// group_id: 群号
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (cli *OnebotApiClientStub) GetGroupInfo(groupId int64, noCache bool) (*model.GroupInfoResult, error) {
	if res, err := cli.Client.GetGroupInfo(context.Background(), &group.GetGroupInfoRequest{
		GroupId: groupId,
		NoCache: noCache,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取群列表
// get_group_list
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (cli *OnebotApiClientStub) GetGroupList(noCache bool) (*model.GroupListResult, error) {
	if res, err := cli.Client.GetGroupList(context.Background(), &wrapperspb.BoolValue{
		Value: noCache,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取群成员信息
// get_group_member_info
// group_id: 群号
// user_id: QQ 号
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (cli *OnebotApiClientStub) GetGroupMemberInfo(groupId int64, userId int64, noCache bool) (*model.GroupMemberInfoResult, error) {
	if res, err := cli.Client.GetGroupMemberInfo(context.Background(), &group.GetGroupMemberInfoRequest{
		GroupId: groupId,
		UserId:  userId,
		NoCache: noCache,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取群成员列表
// get_group_member_list
// group_id: 群号
// no_cache: 是否不使用缓存（使用缓存可能更新不及时，但响应更快）
func (cli *OnebotApiClientStub) GetGroupMemberList(groupId int64, noCache bool) (*model.GroupMemberListResult, error) {
	if res, err := cli.Client.GetGroupMemberList(context.Background(), &group.GetGroupMemberListRequest{
		GroupId: groupId,
		NoCache: noCache,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取群荣誉信息
// get_group_honor_info
// group_id: 群号
// type: 群荣誉类型，目前支持 talkative（群聊之火）、performer（群聊炽焰）、legend（群聊传说）、strong_newbie（群聊新星）、emotion（群表情之火）
func (cli *OnebotApiClientStub) GetGroupHonorInfo(groupId int64, honorType string) (*model.GroupHonorInfoResult, error) {
	if res, err := cli.Client.GetGroupHonorInfo(context.Background(), &group.GetGroupHonorInfoRequest{
		GroupId: groupId,
		Type:    honorType,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取群系统消息
// get_group_system_msg
func (cli *OnebotApiClientStub) GetGroupSystemMsg() (*model.GroupSystemMsgResult, error) {
	if res, err := cli.Client.GetGroupSystemMsg(context.Background(), &emptypb.Empty{}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取精华消息列表
// get_essence_msg_list
// group_id: 群号
func (cli *OnebotApiClientStub) GetEssenceMsgList(groupId int64) (*model.EssenceMsgListResult, error) {
	if res, err := cli.Client.GetEssenceMsgList(context.Background(), &wrapperspb.Int64Value{
		Value: groupId,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取群 @全体成员 剩余次数
// get_group_at_all_remain
// group_id: 群号
func (cli *OnebotApiClientStub) GetGroupAtAllRemain(groupId int64) (*model.GroupAtAllRemainResult, error) {
	if res, err := cli.Client.GetGroupAtAllRemain(context.Background(), &wrapperspb.Int64Value{
		Value: groupId,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 设置群名
// set_group_name
// groupId 群号
func (cli *OnebotApiClientStub) SetGroupName(groupId int64, groupName string) error {
	if _, err := cli.Client.SetGroupName(context.Background(), &groupopt.SetGroupNameRequest{
		GroupId:   groupId,
		GroupName: groupName,
	}); err != nil {
		return err
	}
	return nil
}

// 设置群头像
// set_group_portrait
// groupId 群号
// file 图片文件路径
// cache 缓存时间
func (cli *OnebotApiClientStub) SetGroupPortrait(groupId int64, file string, cache int) error {
	if _, err := cli.Client.SetGroupPortrait(context.Background(), &groupopt.SetGroupPortraitRequest{
		GroupId: groupId,
		File:    file,
		Cache:   int32(cache),
	}); err != nil {
		return err
	}
	return nil
}

// set_group_admin
// 设置群管理员
// groupId 群号
// userId QQ 号
// enable true 为设置，false 为取消
func (cli *OnebotApiClientStub) SetGroupAdmin(groupId int64, userId int64, enable bool) error {
	if _, err := cli.Client.SetGroupAdmin(context.Background(), &groupopt.SetGroupAdminRequest{
		GroupId: groupId,
		UserId:  userId,
		Enable:  enable,
	}); err != nil {
		return err
	}
	return nil
}

// 设置群名片
// set_group_card
// groupId 群号
// userId QQ 号
// card 群名片内容，不填或空字符串表示删除群名片
func (cli *OnebotApiClientStub) SetGroupCard(groupId int64, userId int64, card string) error {
	if _, err := cli.Client.SetGroupCard(context.Background(), &groupopt.SetGroupCardRequest{
		GroupId: groupId,
		UserId:  userId,
		Card:    card,
	}); err != nil {
		return err
	}
	return nil
}

// 设置群头衔
// set_group_special_title
// groupId 群号
// userId QQ 号
// specialTitle 头衔，不填或空字符串表示删除群头衔
// duration 专属头衔有效期, 单位秒, -1 表示永久, 不过此项似乎没有效果, 可能是只有某些特殊的时间长度有效, 有待测试
func (cli *OnebotApiClientStub) SetGroupSpecialTitle(groupId int64, userId int64, specialTitle string, duration uint32) error {
	if _, err := cli.Client.SetGroupSpecialTitle(context.Background(), &groupopt.SetGroupSpecialTitleRequest{
		GroupId:      groupId,
		UserId:       userId,
		SpecialTitle: specialTitle,
		Duration:     duration,
	}); err != nil {
		return err
	}
	return nil
}

// 禁言群成员
// set_group_ban
// groupId 群号
// userId QQ 号
// duration 禁言时长，单位秒，0 表示取消禁言
func (cli *OnebotApiClientStub) SetGroupBan(groupId int64, userId int64, duration uint32) error {
	if _, err := cli.Client.SetGroupBan(context.Background(), &groupopt.SetGroupBanRequest{
		GroupId:  groupId,
		UserId:   userId,
		Duration: duration,
	}); err != nil {
		return err
	}
	return nil
}

// 设置全群禁言
// set_group_whole_ban
// groupId 群号
// enable 是否禁言
func (cli *OnebotApiClientStub) SetGroupWholeBan(groupId int64, enable bool) error {
	if _, err := cli.Client.SetGroupWholeBan(context.Background(), &groupopt.SetGroupWholeBanequest{
		GroupId: groupId,
		Enable:  enable,
	}); err != nil {
		return err
	}
	return nil
}

// 禁言群匿名成员
// set_group_anonymous_ban
// groupId 群号
// anonymous 可选, 要禁言的匿名用户对象（群消息上报的 anonymous 字段）
// anonymousFlag 可选, 要禁言的匿名用户的 flag（需从群消息上报的数据中获得）
// 上面的 anonymous 和 anonymous_flag 两者任选其一传入即可, 若都传入, 则使用 anonymous。
// duration 禁言时长，单位秒，不能超过 30 天
func (cli *OnebotApiClientStub) SetGroupAnonymousBan(groupId int64, anonymous *model.Anonymous, anonymousFlag string, duration uint32) error {
	req := &groupopt.SetGroupAnonymousBanRequest{
		GroupId: groupId,
		// Anonymous:     anonymous.(*model.Anonymous).ToGRPC(),
		AnonymousFlag: anonymousFlag,
		Duration:      duration,
	}
	if anonymous != nil {
		req.Anonymous = anonymous.ToGRPC()
	}
	if _, err := cli.Client.SetGroupAnonymousBan(context.Background(), req); err != nil {
		return err
	}
	return nil
}

// 设置精华消息
// set_essence_msg
// message_id 消息 ID
func (cli *OnebotApiClientStub) SetEssenceMsg(messageId int64) error {
	if _, err := cli.Client.SetEssenceMsg(context.Background(), &wrapperspb.Int64Value{
		Value: messageId,
	}); err != nil {
		return err
	}
	return nil
}

// 删除精华消息
// delete_essence_msg
// message_id 消息 ID
func (cli *OnebotApiClientStub) DeleteEssenceMsg(messageId int64) error {
	if _, err := cli.Client.DeleteEssenceMsg(context.Background(), &wrapperspb.Int64Value{
		Value: messageId,
	}); err != nil {
		return err
	}
	return nil
}

// 发送群签到
// send_group_sign
// groupId 群号
func (cli *OnebotApiClientStub) SendGroupSign(groupId int64) error {
	if _, err := cli.Client.SendGroupSign(context.Background(), &wrapperspb.Int64Value{
		Value: groupId,
	}); err != nil {
		return err
	}
	return nil
}

// 设置群匿名
// set_group_anonymous
// groupId 群号
// enable 是否允许匿名聊天
func (cli *OnebotApiClientStub) SetGroupAnonymous(groupId int64, enable bool) error {
	if _, err := cli.Client.SetGroupAnonymous(context.Background(), &groupopt.SetGroupAnonymousRequest{
		GroupId: groupId,
		Enable:  enable,
	}); err != nil {
		return err
	}
	return nil
}

// 发送群公告
// _send_group_notice
// groupId 群号
// content 公告内容
// image 图片文件路径（可选）
func (cli *OnebotApiClientStub) SendGroupNotice(groupId int64, content string, image string) error {
	if _, err := cli.Client.SendGroupNotice(context.Background(), &groupopt.SendGroupNoticeRequest{
		GroupId: groupId,
		Content: content,
		Image:   image,
	}); err != nil {
		return err
	}
	return nil
}

// 获取群公告
// _get_group_notice
// groupId 群号
func (cli *OnebotApiClientStub) GetGroupNotice(groupId int64) (*model.GroupNoticeResult, error) {
	if res, err := cli.Client.GetGroupNotice(context.Background(), &wrapperspb.Int64Value{
		Value: groupId,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 群组踢人
// set_group_kick
// groupId 群号
// userId QQ 号
// rejectAddRequest 是否拒绝此人的加群请求
func (cli *OnebotApiClientStub) SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) error {
	if _, err := cli.Client.SetGroupKick(context.Background(), &groupopt.SetGroupKickRequest{
		GroupId:          groupId,
		UserId:           userId,
		RejectAddRequest: rejectAddRequest,
	}); err != nil {
		return err
	}
	return nil
}

// 退出群组
// set_group_leave
// groupId 群号
// isDismiss 是否解散，如果登录号是群主，则仅在此项为 true 时能够解散
func (cli *OnebotApiClientStub) SetGroupLeave(groupId int64, isDismiss bool) error {
	if _, err := cli.Client.SetGroupLeave(context.Background(), &groupopt.SetGroupLeaveRequest{
		GroupId:   groupId,
		IsDismiss: isDismiss,
	}); err != nil {
		return err
	}
	return nil
}

// 获取图片
// get_image
func (cli *OnebotApiClientStub) GetImage(file string) (*model.ImageResult, error) {
	if res, err := cli.Client.GetImage(context.Background(), &wrapperspb.StringValue{
		Value: file,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 检查是否可以发送图片
// can_send_image
func (cli *OnebotApiClientStub) CanSendImage() (*model.BoolYesOfResult, error) {
	if res, err := cli.Client.CanSendImage(context.Background(), &emptypb.Empty{}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 图片ORC识别
// ocr_image
func (cli *OnebotApiClientStub) OcrImage(image string) (*model.OcrImageResult, error) {
	if res, err := cli.Client.OcrImage(context.Background(), &wrapperspb.StringValue{
		Value: image,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 发送私信
// send_private_msg
// msg 消息
func (cli *OnebotApiClientStub) SendPrivateMsg(msg *model.PrivateMsg) (*model.SendMessageResult, error) {
	if msg == nil {
		return nil, errors.New("msg is nil")
	}
	if res, err := cli.Client.SendPrivateMsg(context.Background(), msg.ToGRPC()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 发送群消息
// send_group_msg
// msg 消息
func (cli *OnebotApiClientStub) SendGroupMsg(msg *model.GroupMsg) (*model.SendMessageResult, error) {
	if msg == nil {
		return nil, errors.New("msg is nil")
	}
	if res, err := cli.Client.SendGroupMsg(context.Background(), msg.ToGRPC()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 发送消息
func (cli *OnebotApiClientStub) SendMsg(msg *model.MsgForSend) (*model.SendMessageResult, error) {
	if msg == nil {
		return nil, errors.New("msg is nil")
	}
	if res, err := cli.Client.SendMsg(context.Background(), msg.ToGRPC()); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取消息
// get_msg
func (cli *OnebotApiClientStub) GetMsg(id int64) (*model.MessageDataResult, error) {
	if res, err := cli.Client.GetMsg(context.Background(), &wrapperspb.Int64Value{
		Value: id,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 撤回消息
// delete_msg
func (cli *OnebotApiClientStub) DeleteMsg(id int64) error {
	if _, err := cli.Client.DeleteMsg(context.Background(), &wrapperspb.Int64Value{
		Value: id,
	}); err != nil {
		return err
	}
	return nil
}

// 标记消息已读
// mark_msg_as_read
func (cli *OnebotApiClientStub) MarkMsgAsRead(id int64) error {
	if _, err := cli.Client.MarkMsgAsRead(context.Background(), &wrapperspb.Int64Value{
		Value: id,
	}); err != nil {
		return err
	}
	return nil
}

// 获取合并转发消息
// get_forward_msg
func (cli *OnebotApiClientStub) GetForwardMsg(id int64) (*model.ForwardMessageDataResult, error) {
	if res, err := cli.Client.GetForwardMsg(context.Background(), &wrapperspb.Int64Value{
		Value: id,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 发送合并转发 ( 群聊 )
// send_group_forward_msg
func (cli *OnebotApiClientStub) SendGroupForwardMsg(groupId int64, messages []*model.MessageSegment) (*model.SendForwardMessageDataResult, error) {
	if res, err := cli.Client.SendGroupForwardMsg(context.Background(), &message.SendGroupForwardMsgRequest{
		GroupId:  groupId,
		Messages: model.MessageSegmentArray2MessageSegmentGRPCArray(messages),
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 发送合并转发 ( 好友 )
// send_private_forward_msg
func (cli *OnebotApiClientStub) SendPrivateForwardMsg(userId int64, messages []*model.MessageSegment) (*model.SendForwardMessageDataResult, error) {
	if res, err := cli.Client.SendPrivateForwardMsg(context.Background(), &message.SendPrivateForwardMsgGRPC{
		UserId:   userId,
		Messages: model.MessageSegmentArray2MessageSegmentGRPCArray(messages),
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取群消息历史记录
// get_group_msg_history
func (cli *OnebotApiClientStub) GetGroupMsgHistory(message_seq int64, groupId int64) (*model.MessagesResult, error) {
	if res, err := cli.Client.GetGroupMsgHistory(context.Background(), &message.GetGroupMsgHistoryRequest{
		MessageSeq: message_seq,
		GroupId:    groupId,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 获取语音
// get_record
func (cli *OnebotApiClientStub) GetRecord(file string, outFormat string) (*model.RecordResult, error) {
	if res, err := cli.Client.GetRecord(context.Background(), &record.GetRecordRequest{
		File:      file,
		OutFormat: outFormat,
	}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 检查是否可以发送语音
// can_send_record
func (cli *OnebotApiClientStub) CanSendRecord() (*model.BoolYesOfResult, error) {
	if res, err := cli.Client.CanSendRecord(context.Background(), &emptypb.Empty{}); err != nil {
		return nil, err
	} else if res != nil {
		return res.ToStruct(), nil
	}
	return nil, nil
}

// 处理加好友请求
// set_friend_add_request
// flag: 加好友请求的 flag（需从上报的数据中获得）
// approve: 是否同意请求
// remark: 添加后的好友备注（仅在同意时有效）
func (cli *OnebotApiClientStub) SetFriendAddRequest(flag string, approve bool, remark string) error {
	if _, err := cli.Client.SetFriendAddRequest(context.Background(), &request.SetFriendAddRequestRequest{
		Flag:    flag,
		Approve: approve,
		Remark:  remark,
	}); err != nil {
		return err
	}
	return nil
}

// 处理加群请求／邀请
// set_group_add_request
// flag: 加群请求的 flag（需从上报的数据中获得）
// sub_type: add 或 invite，请求类型（需要和上报消息中的 sub_type 字段相符）
// approve: 是否同意请求／邀请
// reason: 拒绝理由（仅在拒绝时有效）
func (cli *OnebotApiClientStub) SetGroupAddRequest(flag string, subType string, approve bool, reason string) error {
	if _, err := cli.Client.SetGroupAddRequest(context.Background(), &request.SetGroupAddRequestRequest{
		Flag:    flag,
		SubType: subType,
		Approve: approve,
		Reason:  reason,
	}); err != nil {
		return err
	}
	return nil
}
