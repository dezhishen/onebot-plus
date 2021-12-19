package main

import (
	"context"
	"time"

	"github.com/dezhishen/onebot-sdk/pkg/event"
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
	return nil
}

func enableHeartbeat(ctx context.Context) error {
	return nil
}
