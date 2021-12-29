package common

import (
	"context"

	"github.com/dezhishen/onebot-plus/example/grpc/proto"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type DemoGRPC interface {
	SayHi(name string) (string, error)
}

//定义插件
type DemoGRPCPlugin struct {
	plugin.Plugin
	// Impl Injection
	Impl DemoGRPC
}

//插件实现GRPC的接口
func (p *DemoGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterDemoServer(s, &DemoGRPCServerStub{Impl: p.Impl})
	return nil
}

type DemoGRPCServerStub struct {
	// This is the real implementation
	Impl DemoGRPC
}

func (m *DemoGRPCServerStub) SayHi(ctx context.Context, req *proto.SayHiReq) (*proto.StringValue, error) {
	v, e := m.Impl.SayHi(req.GetName())
	return &proto.StringValue{Value: v}, e
}

// 业务接口KV的实现，通过gRPC客户端转发请求给插件进程
type DemoServiceClientStub struct {
	client proto.DemoClient
}

//业务接口
func (m *DemoServiceClientStub) SayHi(name string) (string, error) {
	// 转发
	resp, err := m.client.SayHi(context.Background(), &proto.SayHiReq{Name: name})
	if err != nil {
		return "", err
	}
	return resp.Value, err
}

//插件实现GRPC的接口
func (p *DemoGRPCPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &DemoServiceClientStub{client: proto.NewDemoClient(c)}, nil
}
