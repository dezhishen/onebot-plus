package main

import (
	"context"
	"fmt"
	"net/rpc"
	"os"
	"os/exec"
	"time"

	"github.com/dezhishen/onebot-sdk/pkg/event"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	go func() {
		time.Sleep(time.Second * 5)
		//模拟退出
		log.Info("模拟退出...")
		cancel()
	}()
	err := loadingPlugin(ctx)
	if err != nil {
		log.Errorf("加载插件失败:%v", err)
		cancel()
	}
	defer cancel()
	go func() {
		err := startWs(ctx)
		if err != nil {
			log.Errorf("监听事件发生错误:%v", err)
			cancel()
		}
		//todo 如果健康检查开启，则开启心跳
		err = enableHeartbeat(ctx)
		if err != nil {
			log.Errorf("开启心跳失败:%v", err)
			cancel()
		}
	}()
	//dosomething else
	<-ctx.Done()
}

func startWs(ctx context.Context) error {
	return event.StartWsWithContext(ctx)
}

func loadingPlugin(ctx context.Context) error {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugin/greeter"),
		Logger:          logger,
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter := raw.(Greeter)
	fmt.Println(greeter.Greet())
	return nil
}

func enableHeartbeat(ctx context.Context) error {
	return nil
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"greeter": &GreeterPlugin{},
}

type Greeter interface {
	Greet() string
}

type GreeterPlugin struct {
	// Impl Injection
	Impl Greeter
}

type GreeterRPC struct{ client *rpc.Client }
type GreeterRPCServer struct {
	// This is the real implementation
	Impl Greeter
}

func (p *GreeterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{Impl: p.Impl}, nil
}

func (GreeterPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: c}, nil
}
