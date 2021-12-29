package common

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Demo interface {
	SayHi() string
}

type DemoRPC struct {
	client *rpc.Client
}

func (d *DemoRPC) SayHi() string {
	var resp string
	err := d.client.Call("Plugin.SayHi", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

type DemoRPCServer struct {
	Impl Demo
}

func (s *DemoRPCServer) SayHi(args interface{}, resp *string) error {
	*resp = s.Impl.SayHi()
	return nil
}

type DemoPlugin struct {
	// Impl Injection
	Impl Demo
}

func (p *DemoPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &DemoRPCServer{Impl: p.Impl}, nil
}

func (DemoPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &DemoRPC{client: c}, nil
}
