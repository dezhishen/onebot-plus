# onebot-plus
onebot拓展平台
# 项目目标
本项目意在打造一个`使用简单`，`拓展简单`的`插件化`的onebot能力拓展平台
* 基于`onebot协议`的QQ机器人的`外部拓展`，即机器人与项目分别部署
* 插件基于`grpc`与主程序隔离
* 使用[`dezhishen/onebot-sdk`](https://github.com/dezhishen/onebot-sdk)作为sdk
* 管理方法
    - [ ] 基于QQ命令的管理
    - [ ] 基于web-api的管理
    - [ ] 拓展WEB-UI管理界面
* 使用docker部署，降低部署和运维难度
# 快速开始
## 二进制
- 下载[release](https://github.com/dezhishen/onebot-plus/releases)中的二进制文件
- 在当前目录下创建`configs/onebot-plus.yaml`文件
  - 参考[配置文件](./configs/onebot-plus.yaml)
- 在当前目录下创建`configs/plugins`文件夹
  - 将插件放入该文件夹
# 插件和二次开发指引
- [Wiki](./wiki/develop)
