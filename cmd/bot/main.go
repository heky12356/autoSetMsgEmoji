package main

import (
	"autoSetMsgEmoji/internal/config"
	"autoSetMsgEmoji/internal/connect"
)

func main() {
	config.Init()
	c, err := connect.Init(config.Config.Hostadd)
	if err != nil {
		panic(err)
	}
	defer c.Close()
}
