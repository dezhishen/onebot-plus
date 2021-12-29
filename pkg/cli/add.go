package cli

import (
	context "context"

	"github.com/sirupsen/logrus"
)

type AddHelper interface {
	Sum(int64, int64) (int64, error)
}

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCAddHelperClient struct{ Client AddHelperClient }

func (m *GRPCAddHelperClient) Sum(a, b int64) (int64, error) {
	resp, err := m.Client.Sum(context.Background(), &SumRequest{
		A: a,
		B: b,
	})
	if err != nil {
		logrus.Info("add.Sum", "client", "start", "err", err)
		return 0, err
	}
	return resp.R, err
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCAddHelperServer struct {
	// This is the real implementation
	Impl AddHelper
}

func (m *GRPCAddHelperServer) Sum(ctx context.Context, req *SumRequest) (resp *SumResponse, err error) {
	r, err := m.Impl.Sum(req.A, req.B)
	if err != nil {
		return nil, err
	}
	return &SumResponse{R: r}, err
}
