package plugin

import (
	"log"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

// // 插件的接口
// type BaseInterface interface {
// 	// 插件ID
// 	ID() string
// 	// 插件名称
// 	Name() string
// 	// 插件描述
// 	Description() string
// 	// 插件使用帮助
// 	Help() string
// 	// 插件设置项目
// 	AllSettings() []OnebotPlusPluginSetting
// }

// type BaseRPC struct {
// 	client *rpc.Client
// }

// type BaseRPCServer struct {
// 	Impl BaseInterface
// }

// // 插件ID
// func (p BaseRPC) ID() string {
// 	var resp string
// 	err := p.client.Call("Plugin.ID", new(interface{}), &resp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return resp
// }

// // 插件名称
// func (p BaseRPC) Name() string {
// 	var resp string
// 	err := p.client.Call("Plugin.Name", new(interface{}), &resp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return resp
// }

// // 插件描述
// func (p BaseRPC) Description() string {
// 	var resp string
// 	err := p.client.Call("Plugin.Description", new(interface{}), &resp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return resp
// }

// // 插件使用帮助
// func (p BaseRPC) Help() string {
// 	var resp string
// 	err := p.client.Call("Plugin.Help", new(interface{}), &resp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return resp
// }

// // 插件设置项目
// func (p BaseRPC) AllSettings() []OnebotPlusPluginSetting {
// 	var resp []OnebotPlusPluginSetting
// 	err := p.client.Call("Plugin.AllSettings", new(interface{}), &resp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return resp
// }

func LoadPlugin(path string) PrivateMssageHanlder {
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"main": &PrivateMessageGRPCPlugin{},
		},
		Cmd: exec.Command(path),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	// defer client.Kill()
	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}
	// Request the plugin
	raw, err := rpcClient.Dispense("main")
	if err != nil {
		log.Fatal(err)
	}
	d := raw.(PrivateMssageHanlder)
	return d
}
