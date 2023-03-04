package base

import (
	context "context"

	"github.com/dezhishen/onebot-plus/pkg/plugin/api"
	"github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type OnebotPluginBaseServerStub struct {
	Broker *plugin.GRPCBroker
	Impl   OnebotPluginBase
}

func (s *OnebotPluginBaseServerStub) Id(ctx context.Context, in *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return &wrapperspb.StringValue{Value: s.Impl.Id()}, nil
}

func (s *OnebotPluginBaseServerStub) Name(ctx context.Context, in *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return &wrapperspb.StringValue{Value: s.Impl.Name()}, nil
}

func (s *OnebotPluginBaseServerStub) Description(ctx context.Context, in *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return &wrapperspb.StringValue{Value: s.Impl.Description()}, nil
}

func (s *OnebotPluginBaseServerStub) Help(ctx context.Context, in *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return &wrapperspb.StringValue{Value: s.Impl.Help()}, nil
}

func (s *OnebotPluginBaseServerStub) Init(ctx context.Context, in *wrapperspb.UInt32Value) (*emptypb.Empty, error) {
	conn, err := s.Broker.Dial(in.Value)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	// defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := s.Impl.Init(client)
	return &emptypb.Empty{}, e
}

func (s *OnebotPluginBaseServerStub) BeforeExit(ctx context.Context, in *wrapperspb.UInt32Value) (*emptypb.Empty, error) {
	conn, err := s.Broker.Dial(in.Value)
	if err != nil {
		logrus.Errorf("conn err %v", err)
		return nil, err
	}
	defer conn.Close()
	client := &api.OnebotApiClientStub{
		Client: api.NewOnebotAPiClientInterfaceGRPCServiceClient(conn),
	}
	e := s.Impl.BeforeExit(client)
	return &emptypb.Empty{}, e
}
