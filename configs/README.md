# configs
配置文件读取路径为程序运行目录下的`configs`文件夹下的`onebot-plus.yaml`文件
## onebot-plus.yaml
```
api: 
# 调用go-cqhttp的api 配置
  type: http 
  #支持三种方式 http / websocket / websocket-reverse
  endpoint: "http(s)://xxxxxxxx" 
  # 当类型是websocket-reverse时，此项用于声明本服务gin的启动配置例如:`0.0.0.0:6700`或者 `127.0.0.1:6700`
  accessToken: "_access_token"
event:
  type: "websocket" 
  # websocket / websocket-reverse / http-reverse
  addr: "ws(s)://xxxxxx" 
  #当类型是websocket-reverse/http-reverse时，此项用于声明本服务gin的启动配置例如:`0.0.0.0:6700`或者 `127.0.0.1:6700`
  accessToken: "_access_token"
```
## 注意事项
如果是反向ws或者反向http，需要在go-cqhttp的配置文件中修改配置,用于连接到本服务