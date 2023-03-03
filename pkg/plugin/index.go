package plugin

import (
	"github.com/dezhishen/onebot-plus/pkg/plugin/base"
	"github.com/dezhishen/onebot-plus/pkg/plugin/event"
)

type OnebotPlugin interface {
	base.OnebotPluginBase
	event.OnebotEventCallBackInterface
}
