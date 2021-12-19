# onebot-plus
onebot拓展平台
# 项目目标
本项目意在打造一个`使用简单`，`拓展简单`的`插件化`的onebot能力拓展平台
* 基于`onebot协议`的QQ机器人的`外部拓展`，即机器人与项目分别部署
* 使用[`dezhishen/onebot-sdk`](https://github.com/dezhishen/onebot-sdk)作为sdk
* 基于[`hashicorp/go-plugin`](https://github.com/hashicorp/go-plugin)作为插件开发核心库
* 管理方法
    * 基于QQ命令的管理
    * 基于web-api的管理
    * 拓展WEB-UI管理界面
* 使用docker部署，降低部署和运维难度
# 功能清单
**本项目仍在开发阶段，功能设计和实现可能会有变化**
## 平台-onebot
- [ ] 引入[`dezhishen/onebot-sdk`](https://github.com/dezhishen/onebot-sdk)
- [ ] 基于WS心跳检测的健康检查
- [ ] onebot管理
## 平台
- [ ] 基础配置
- [ ] 平台生命周期管理
- [ ] 平台指令设计
## 平台-插件
- [ ] 引入[`hashicorp/go-plugin`](https://github.com/hashicorp/go-plugin)
- [ ] 设计事件下发与onebot交互上报
- [ ] 事件下发
- [ ] 插件调度onebot

