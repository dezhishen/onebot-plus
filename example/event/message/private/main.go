package main

import (
	"io/ioutil"
	"log"

	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func main() {
	// We don't want to see the plugin logs.
	log.SetOutput(ioutil.Discard)
	path := "./impl/plugin.exe"
	d := plugin.LoadPlugin(path)
	msg := &model.EventMessagePrivate{}
	msg.Sender = &model.Sender{
		Nickname: "test",
	}
	d.Handle(msg)
}
