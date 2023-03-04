package cmd

import (
	"context"

	"github.com/dezhishen/onebot-plus/pkg/pluginmanager"
)

func StartWithContext(ctx context.Context) error {
	// onebotCli := &cli.OnebotCliRealImpl{}
	return pluginmanager.Start(ctx)
}

// 启动
func Start() error {
	ctx := context.Background()
	return pluginmanager.Start(ctx)
}
