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
	defer cancel()
	go func() {
		err := startWs(ctx)
		if err != nil {
			log.Errorf("监听事件发生错误:%v", err)
			cancel()
		}
	}()
	<-ctx.Done()
}

func startWs(ctx context.Context) error {
	return event.StartWsWithContext(ctx)
}
