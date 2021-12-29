package main

import "github.com/dezhishen/onebot-plus/pkg/plugin"

func main() {
	plugin.NewOnebotEventPluginBuilder().
		//设置插件内容
		Id("test").Name("测试插件").Description("这是测试插件").
		//构建插件
		Build().
		//启动
		Start()
}
