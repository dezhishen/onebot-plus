package base

import (
	context "context"

	"github.com/dezhishen/onebot-plus/pkg/plugin/api"
	sdk_api "github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type OnebotPluginBaseClientStub struct {
	Broker *plugin.GRPCBroker
	Client OnebotPluginBaseGRPCClient
}

// 插件Id
func (c *OnebotPluginBaseClientStub) Id() string {
	stringValue, err := c.Client.Id(context.Background(), &emptypb.Empty{})
	if err != nil {
		logrus.Errorf("Id() err %v", err)
		return ""
	}
	return stringValue.Value
}

// 插件名称
func (c *OnebotPluginBaseClientStub) Name() string {
	stringValue, err := c.Client.Name(context.Background(), &emptypb.Empty{})
	if err != nil {
		logrus.Errorf("Name() err %v", err)
		return ""
	}
	return stringValue.Value
}

// 插件描述
func (c *OnebotPluginBaseClientStub) Description() string {
	stringValue, err := c.Client.Description(context.Background(), &emptypb.Empty{})
	if err != nil {
		logrus.Errorf("Description() err %v", err)
		return ""
	}
	return stringValue.Value
}

// 插件帮助
func (c *OnebotPluginBaseClientStub) Help() string {
	stringValue, err := c.Client.Help(context.Background(), &emptypb.Empty{})
	if err != nil {
		logrus.Errorf("Help() err %v", err)
		return ""
	}
	return stringValue.Value
}

// 生命周期
func (c *OnebotPluginBaseClientStub) Init(onebotApi sdk_api.OnebotApiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		api.RegisterOnebotApiGRPCServiceServer(s,
			&api.OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := c.Broker.NextId()
	go c.Broker.AcceptAndServe(brokerID, serverFunc)
	_, err := c.Client.Init(context.Background(), &wrapperspb.UInt32Value{
		Value: brokerID,
	})
	if err != nil {
		return err
	}
	// s.Stop()
	return nil
}

// 退出前
func (c *OnebotPluginBaseClientStub) BeforeExit(onebotApi sdk_api.OnebotApiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		api.RegisterOnebotApiGRPCServiceServer(s,
			&api.OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := c.Broker.NextId()
	go c.Broker.AcceptAndServe(brokerID, serverFunc)
	_, err := c.Client.BeforeExit(context.Background(), &wrapperspb.UInt32Value{
		Value: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}
