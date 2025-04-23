package connect

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func Init(host string) (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: host, Path: "/"}
	log.Printf("连接到 %s", u.String())

	// 建立 WebSocket 连接
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("连接失败:", err)
	}
	return c, nil
}
