package api

import (
	"context"

	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/api/account"
	"github.com/dezhishen/onebot-sdk/pkg/api/cqhttp"
	"github.com/dezhishen/onebot-sdk/pkg/api/file"
	"github.com/dezhishen/onebot-sdk/pkg/api/friend"
	"github.com/dezhishen/onebot-sdk/pkg/api/friendopt"
	"github.com/dezhishen/onebot-sdk/pkg/api/group"
	"github.com/dezhishen/onebot-sdk/pkg/api/groupopt"
	"github.com/dezhishen/onebot-sdk/pkg/api/image"
	"github.com/dezhishen/onebot-sdk/pkg/api/message"
	"github.com/dezhishen/onebot-sdk/pkg/api/record"
	"github.com/dezhishen/onebot-sdk/pkg/api/request"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type OnebotApiGRPCPlugin struct {
	// 需要嵌入插件接口
	plugin.Plugin
	// 具体实现，仅当业务接口实现基于Go时该字段有用
	Impl api.OnebotApiClientInterface
}

// 插件实现GRPC的接口
func (p *OnebotApiGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	svcStub := &OnebotApiServerStub{Impl: p.Impl}
	account.RegisterOnebotApiAccountGRPCServiceServer(s, svcStub)
	cqhttp.RegisterOnebotApiCqhttpGRPCServiceServer(s, svcStub)
	file.RegisterOnebotApiFileGRPCServiceServer(s, svcStub)
	friend.RegisterOnebotApiFriendGRPCServiceServer(s, svcStub)
	friendopt.RegisterOnebotApiFriendOptGRPCServiceServer(s, svcStub)
	group.RegisterOnebotApiGroupGRPCServiceServer(s, svcStub)
	groupopt.RegisterOnebotApiGroupOptGRPCServiceServer(s, svcStub)
	image.RegisterOnebotApiImageGRPCServiceServer(s, svcStub)
	message.RegisterOnebotApiMessageGRPCServiceServer(s, svcStub)
	record.RegisterOnebotApiRecordGRPCServiceServer(s, svcStub)
	request.RegisterOnebotApiRequestGRPCServiceServer(s, svcStub)
	return nil
}

// 插件实现GRPC的接口
func (p *OnebotApiGRPCPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &OnebotApiClientStub{Client: NewOnebotAPiClientInterfaceGRPCServiceClient(
		c,
	)}, nil
}
