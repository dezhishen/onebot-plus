package main

import (
	"os"

	"github.com/dezhishen/onebot-plus/example/rpc/common"
	"github.com/dezhishen/onebot-plus/pkg/plugins"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type DemoRelImp struct {
	logger hclog.Logger
}

func (d *DemoRelImp) SayHi() string {
	d.logger.Debug("message from DemoPlugin")
	return "Hello!"
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	demo := &DemoRelImp{
		logger: logger,
	}
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"demo": &common.DemoPlugin{Impl: demo},
	}

	logger.Debug("message from plugin", "sayhello")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugins.HandshakeConfig,
		Plugins:         pluginMap,
	})
}
