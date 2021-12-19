package plugins

type GrpcPlugin interface {
}

// 插件的接口
type OnebotPlusPluginInterface interface {
	GrpcPlugin
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
