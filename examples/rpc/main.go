package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/dezhishen/onebot-plus/example/rpc/common"
	onebotPlugin "github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func main() {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: onebotPlugin.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"demo": &common.DemoPlugin{},
		},
		Cmd:    exec.Command("./demo_impl.exe"),
		Logger: logger,
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}
	// Request the plugin
	raw, err := rpcClient.Dispense("demo")
	if err != nil {
		log.Fatal(err)
	}
	greeter := raw.(common.Demo)
	fmt.Println(greeter.SayHi())
	fmt.Println(greeter.SayHi())
}
