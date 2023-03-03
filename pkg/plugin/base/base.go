package base

import "github.com/dezhishen/onebot-sdk/pkg/api"

type OnebotPluginBase interface {
	// 插件Id
	Id() string
	// 插件名称
	Name() string
	// 插件描述
	Description() string
	// 插件帮助
	Help() string
	// 生命周期
	Init(cli api.OnebotAPiClientInterface) error
}
