package plugins

import (
	"net/rpc"
)

// 插件的接口
type BaseInterface interface {
	// 插件ID
	ID() string
	// 插件名称
	Name() string
	// 插件描述
	Description() string
	// 插件使用帮助
	Help() string
	// 插件设置项目
	AllSettings() []OnebotPlusPluginSetting
}

type BaseRPC struct {
	client *rpc.Client
}

type BaseRPCServer struct {
	Impl BaseInterface
}

// 插件ID
func (p BaseRPC) ID() string {
	var resp string
	err := p.client.Call("Plugin.ID", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

// 插件名称
func (p BaseRPC) Name() string {
	var resp string
	err := p.client.Call("Plugin.Name", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

// 插件描述
func (p BaseRPC) Description() string {
	var resp string
	err := p.client.Call("Plugin.Description", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

// 插件使用帮助
func (p BaseRPC) Help() string {
	var resp string
	err := p.client.Call("Plugin.Help", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

// 插件设置项目
func (p BaseRPC) AllSettings() []OnebotPlusPluginSetting {
	var resp []OnebotPlusPluginSetting
	err := p.client.Call("Plugin.AllSettings", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}
