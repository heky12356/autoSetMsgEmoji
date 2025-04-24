package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"autoSetMsgEmoji/internal/config"
	"autoSetMsgEmoji/internal/connect"
	"autoSetMsgEmoji/internal/model"
	"autoSetMsgEmoji/internal/service"

	"github.com/gorilla/websocket"
)

func main() {
	config.Init()
	c, err := connect.Init(config.Config.Hostadd)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// 监听中断信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	done := make(chan struct{})

	// 启动协程读取消息
	go func() {
		defer close(done)
		for {
			t, message, err := c.ReadMessage()
			if err != nil {
				log.Println("读取消息失败:", err)
				return
			}
			log.Printf("接收到消息: %s", message)
			var msg model.Response
			err = json.Unmarshal(message, &msg)
			if err != nil {
				log.Println("消息反序列化失败:", err)
			}
			log.Print(msg)
			if msg.User_id == config.Config.TargetId && msg.Message_type == "group" {
				err = service.SendGroupMsgEmoji(c, msg.Message_id)
				if err != nil {
					log.Println(err)
				}
			}

			log.Println(t)
		}
	}()

	// 主循环，处理中断信号
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("接收到中断信号，关闭连接")
			// 发送关闭消息
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("发送关闭消息失败:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
