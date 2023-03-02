package plugin

import (
	"context"

	"github.com/dezhishen/onebot-sdk/pkg/api"
	_ "github.com/dezhishen/onebot-sdk/pkg/event"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// 业务接口的实现，通过gRPC客户端转发请求给插件进程
type OnebotEventClientStub struct {
	broker *plugin.GRPCBroker
	client EventCallbackGRPCClient
	// Client EventCallbackGRPCServer
}

func (eCli *OnebotEventClientStub) HandleMessagePrivate(data *model.EventMessagePrivate, onebotApi api.OnebotAPiClientInterface) error {
	oneApiServer := &OnebotApiServerStub{
		Impl: onebotApi,
	}
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s, oneApiServer)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleMessagePrivate(context.Background(), &EventMessagePrivateGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleMessageGroup(data *model.EventMessageGroup, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleMessageGroup(context.Background(), &EventMessageGroupGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleNoticeGroupUpload(data *model.EventNoticeGroupUpload, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleNoticeGroupUpload(context.Background(), &EventNoticeGroupUploadGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleNoticeGroupAdmin(data *model.EventNoticeGroupAdmin, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleNoticeGroupAdmin(context.Background(), &EventNoticeGroupAdminGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleNoticeGroupDecrease(data *model.EventNoticeGroupDecrease, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleNoticeGroupDecrease(context.Background(), &EventNoticeGroupDecreaseGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleNoticeGroupIncrease(data *model.EventNoticeGroupIncrease, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleNoticeGroupIncrease(context.Background(), &EventNoticeGroupIncreaseGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleNoticeGroupBan(data *model.EventNoticeGroupBan, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleNoticeGroupBan(context.Background(), &EventNoticeGroupBanGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleNoticeFriendAdd(data *model.EventNoticeFriendAdd, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleNoticeFriendAdd(context.Background(), &EventNoticeFriendAddGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil

}

func (eCli *OnebotEventClientStub) HandleNoticeGroupRecall(data *model.EventNoticeGroupRecall, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleNoticeGroupRecall(context.Background(), &EventNoticeGroupRecallGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleNoticeFriendRecall(data *model.EventNoticeFriendRecall, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleNoticeFriendRecall(context.Background(), &EventNoticeFriendRecallGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil

}

func (eCli *OnebotEventClientStub) HandleNoticeGroupNotifyPoke(data *model.EventNoticeGroupNotifyPoke, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleNoticeGroupNotifyPoke(context.Background(), &EventNoticeGroupNotifyPokeGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleNoticeGroupNotifyHonor(data *model.EventNoticeGroupNotifyHonor, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleNoticeGroupNotifyHonor(context.Background(), &EventNoticeGroupNotifyHonorGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleNoticeGroupNotifyLuckyKing(data *model.EventNoticeGroupNotifyLuckyKing, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleNoticeGroupNotifyLuckyKing(context.Background(), &EventNoticeGroupNotifyLuckyKingGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleRequestFriend(data *model.EventRequestFriend, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleRequestFriend(context.Background(), &EventRequestFriendGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleRequestGroup(data *model.EventRequestGroup, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleRequestGroup(context.Background(), &EventRequestGroupGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleMetaLifecycle(data *model.EventMetaLifecycle, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleMetaLifecycle(context.Background(), &EventMetaLifecycleGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}

func (eCli *OnebotEventClientStub) HandleMetaHeartBeat(data *model.EventMetaHeartbeat, onebotApi api.OnebotAPiClientInterface) error {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		var finalOpts []grpc.ServerOption
		finalOpts = append(finalOpts, opts...)
		finalOpts = append(finalOpts, grpc.MaxRecvMsgSize(256*1024*1024), grpc.MaxSendMsgSize(256*1024*1024))
		s = grpc.NewServer(finalOpts...)
		RegisterOnebotApiGRPCServiceServer(s,
			&OnebotApiServerStub{
				Impl: onebotApi,
			},
		)
		return s
	}
	brokerID := eCli.broker.NextId()
	go eCli.broker.AcceptAndServe(brokerID, serverFunc)
	msg := data.ToGRPC()
	_, err := eCli.client.HandleMetaHeartBeat(context.Background(), &EventMetaHeartbeatGRPCWithOnebotApi{
		Message:   msg,
		OnebotApi: brokerID,
	})
	if err != nil {
		return err
	}
	s.Stop()
	return nil
}
