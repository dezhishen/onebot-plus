package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/dezhishen/onebot-plus/example/grpc/common"
	onebotPlugin "github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
)

func main() {
	// We don't want to see the plugin logs.
	log.SetOutput(ioutil.Discard)
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: onebotPlugin.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"demo": &common.DemoGRPCPlugin{},
		},
		Cmd: exec.Command("./plugin/plugin.exe"),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	defer client.Kill()
	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		logrus.Errorf("%v", err)
		log.Fatal(err)
	}
	// Request the plugin
	raw, err := rpcClient.Dispense("demo")
	if err != nil {
		log.Fatal(err)
	}
	demoGRPC := raw.(common.DemoGRPC)
	msg, err := demoGRPC.SayHi("main")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
