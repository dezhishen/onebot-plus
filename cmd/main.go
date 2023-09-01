package main

import "github.com/dezhishen/onebot-plus/pkg/cmd"

func main() {
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}
