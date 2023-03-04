# develop

## 插件说明
插件基于回调函数进行开发
### 配置
* Id("test") //设置ID
* Name("测试插件") //设置名称
* Description("这是测试插件") //设置描述
* Help("插件帮助") //设置帮助文本
### 插件生命周期
* Init(cli cli.Bot)error //初始化 **cli的grpc连接不会断开**
* BeforeExit(cli cli.Bot)error //插件退出前回调
### Onebot事件回调
[https://github.com/botuniverse/onebot-11/tree/master/event](https://github.com/botuniverse/onebot-11/tree/master/event)
## 插件开发
### go语言
本项目提供了一套`plugin.Buidler`用于插件开发

可以定义插件的配置、插件生命周期的回调函数、onebot机器人事件的回调函数
```
package main

import (
	"time"

	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/sirupsen/logrus"
)

func main() {
	plugin.OnebotPluginBuilder().
		//设置插件内容
		Id("test").
		Name("测试插件").
		Description("这是测试插件").
		Help("测试查询,无操作指令").
		Init(func(cli api.OnebotApiClientInterface) error {
			go func() {
				for {
					result, err := cli.GetLoginInfo()
					if err != nil {
						logrus.Errorf("GetLoginInfo err %v", err)
					} else {
						logrus.Infof(result.Data.Nickname)
					}
					time.Sleep(time.Second * 5)
				}
			}()
			return nil
		}).
		HandleMessageGroup(func(data *model.EventMessageGroup, onebotApi api.OnebotApiClientInterface) error {
			var msgs []*model.MessageSegment
			msgs = append(msgs, &model.MessageSegment{
				Type: "text",
				Data: &model.MessageElementText{
					Text: "你好,我复读机\n",
				},
			})
			msg := append(msgs, data.Message...)
			onebotApi.SendGroupMsg(
				&model.GroupMsg{
					GroupId: data.GroupId,
					Message: msg,
				},
			)
			return nil
		}).
		//构建插件
		Build().
		Start()
}

```

### 其他语言
* proto文件
    * [plus_base.proto](../../pkg/plugin/base/plus_base.proto)
    * [plus_event.proto](../../pkg/plugin/event/plus_event.proto)
* 插件开发方式
    * [`hashicorp/go-plugin`](https://github.com/hashicorp/go-plugin)