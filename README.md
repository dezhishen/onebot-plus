# onebot-plus
onebot拓展平台
# 项目目标
本项目意在打造一个`使用简单`，`拓展简单`的`插件化`的onebot能力拓展平台
* 基于`onebot协议`的QQ机器人的`外部拓展`，即机器人与项目分别部署
* 插件基于`grpc`与主程序隔离
* 使用[`dezhishen/onebot-sdk`](https://github.com/dezhishen/onebot-sdk)作为sdk
* 管理方法
    * 基于QQ命令的管理
    * 基于web-api的管理
    * 拓展WEB-UI管理界面
* 使用docker部署，降低部署和运维难度
# 插件和二次开发指引
- [Wiki](./wiki/develop)
# 功能清单
**本项目仍在开发阶段，功能设计和实现可能会有变化**
## 功能
### 核心
- [x] 引入[`dezhishen/onebot-sdk`](https://github.com/dezhishen/onebot-sdk)
- [x] 引入[`hashicorp/go-plugin`](https://github.com/hashicorp/go-plugin)
### 将onebot-sdk中的API写成插件让其他插件可以调用
- [x] `send_private_msg` 发送私聊消息
- [x] `send_group_msg` 发送群消息
- [x] `send_msg` 发送消息
- [x] `delete_msg` 撤回消息
- [x] `get_msg` 获取消息
- [x] `get_forward_msg` 获取合并转发消息
- [x] `set_group_kick` 群组踢人
- [x] `set_group_ban` 群组单人禁言
- [x] `set_group_anonymous_ban` 群组匿名用户禁言
- [x] `set_group_whole_ban` 群组全员禁言
- [x] `set_group_admin` 群组设置管理员
- [x] `set_group_anonymous` 群组匿名
- [x] `set_group_card` 设置群名片（群备注）
- [x] `set_group_name` 设置群名
- [x] `set_group_leave` 退出群组
- [x] `set_group_special_title` 设置群组专属头衔
- [x] `set_group_add_request` 处理加群请求／邀请
- [x] `get_group_info` 获取群信息
- [x] `get_group_list` 获取群列表
- [x] `get_group_member_info` 获取群成员信息
- [x] `get_group_member_list` 获取群成员列表
- [x] `get_group_honor_info` 获取群荣誉信息
- [x] `send_like` 发送好友赞
- [x] `set_friend_add_request` 处理加好友请求
- [x] `get_friend_list` 获取好友列表
- [x] `get_login_info` 获取登录号信息
- [x] `get_stranger_info` 获取陌生人信息
- [x] `get_cookies` 获取 Cookies
- [x] `get_csrf_token` 获取 CSRF Token 
- [x] `get_credentials` 获取 QQ 相关接口凭证
- [x] `get_record` 获取语音
- [x] `get_image` 获取图片
- [x] `get_status` 获取运行状态
- [x] `get_version_info` 获取版本信息
- [x] `set_restart` 重启 OneBot 实现
- [x] `clean_cache` 清理缓存
- [x] `can_send_image` 检查是否可以发送图片
- [x] `can_send_record` 检查是否可以发送语音
### 增加插件生命周期回调函数
- [x] `Init` 初始化
- [x] `BeforeExit` 退出前回调
### 将事件回调写成插件给外部插件调用
- [x] 群聊消息
- [x] 私聊消息
- [x] 群文件上传
- [x] 群管理员变动
- [x] 群成员减少
- [x] 群成员增加
- [x] 群禁言
- [x] 好友添加
- [x] 群消息撤回
- [x] 好友消息撤回
- [x] 群内戳一戳
- [x] 群红包运气王
- [x] 群成员荣誉变更
- [x] 加好友请求
- [x] 加群请求／邀请
- [x] 生命周期
- [x] 心跳
- [x] 相关配置