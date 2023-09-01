package plugin

import (
	"github.com/dezhishen/onebot-plus/pkg/plugin/base"
	"github.com/dezhishen/onebot-plus/pkg/plugin/event"
)

type OnebotClientStub struct {
	base.OnebotPluginBaseClientStub
	event.OnebotEventCallbackClientStub
}
