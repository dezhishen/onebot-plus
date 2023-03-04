package event

import (
	"context"

	"github.com/dezhishen/onebot-plus/pkg/plugin/api"
	sdk_api "github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// 业务接口的实现，通过gRPC客户端转发请求给插件进程
type OnebotEventCallbackClientStub struct {
	Broker *plugin.GRPCBroker
	Client OnebotEventCallbackGRPCClient
}

func (eCli *OnebotEventCallbackClientStub) HandleMessagePrivate(data *model.EventMessagePrivate, onebotApi sdk_api.OnebotApiClientInterface) error {
	oneApiServer := &api.OnebotApiServerStub{
		Impl: onebotApi,
	}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		api.RegisterOnebotApiGRPCServiceServer(s, oneApiServer)
		return s
	}
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleMessagePrivate(context.Background(), &EventMessagePrivateGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleMessageGroup(data *model.EventMessageGroup, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleMessageGroup(context.Background(), &EventMessageGroupGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleNoticeGroupUpload(data *model.EventNoticeGroupUpload, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleNoticeGroupUpload(context.Background(), &EventNoticeGroupUploadGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleNoticeGroupAdmin(data *model.EventNoticeGroupAdmin, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleNoticeGroupAdmin(context.Background(), &EventNoticeGroupAdminGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleNoticeGroupDecrease(data *model.EventNoticeGroupDecrease, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleNoticeGroupDecrease(context.Background(), &EventNoticeGroupDecreaseGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleNoticeGroupIncrease(data *model.EventNoticeGroupIncrease, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleNoticeGroupIncrease(context.Background(), &EventNoticeGroupIncreaseGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleNoticeGroupBan(data *model.EventNoticeGroupBan, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleNoticeGroupBan(context.Background(), &EventNoticeGroupBanGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleNoticeFriendAdd(data *model.EventNoticeFriendAdd, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleNoticeFriendAdd(context.Background(), &EventNoticeFriendAddGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil

}

func (eCli *OnebotEventCallbackClientStub) HandleNoticeGroupRecall(data *model.EventNoticeGroupRecall, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleNoticeGroupRecall(context.Background(), &EventNoticeGroupRecallGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleNoticeFriendRecall(data *model.EventNoticeFriendRecall, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleNoticeFriendRecall(context.Background(), &EventNoticeFriendRecallGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil

}

func (eCli *OnebotEventCallbackClientStub) HandleNoticeGroupNotifyPoke(data *model.EventNoticeGroupNotifyPoke, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleNoticeGroupNotifyPoke(context.Background(), &EventNoticeGroupNotifyPokeGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleNoticeGroupNotifyHonor(data *model.EventNoticeGroupNotifyHonor, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleNoticeGroupNotifyHonor(context.Background(), &EventNoticeGroupNotifyHonorGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleNoticeGroupNotifyLuckyKing(data *model.EventNoticeGroupNotifyLuckyKing, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleNoticeGroupNotifyLuckyKing(context.Background(), &EventNoticeGroupNotifyLuckyKingGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleRequestFriend(data *model.EventRequestFriend, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleRequestFriend(context.Background(), &EventRequestFriendGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleRequestGroup(data *model.EventRequestGroup, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleRequestGroup(context.Background(), &EventRequestGroupGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleMetaLifecycle(data *model.EventMetaLifecycle, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleMetaLifecycle(context.Background(), &EventMetaLifecycleGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventCallbackClientStub) HandleMetaHeartBeat(data *model.EventMetaHeartbeat, onebotApi sdk_api.OnebotApiClientInterface) error {
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
	brokerID := eCli.Broker.NextId()
	go eCli.Broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.Client.HandleMetaHeartBeat(context.Background(), &EventMetaHeartbeatGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}
