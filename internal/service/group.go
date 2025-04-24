package service

import (
	"encoding/json"
	"fmt"

	"autoSetMsgEmoji/internal/model"

	"github.com/gorilla/websocket"
)

func SendGroupMsg(c *websocket.Conn, GroupId int64, msg string) (err error) {
	// fmt.Println("GroupId:", GroupId, "msg:", msg)
	ws_msg := model.Message{
		Action: "send_group_msg",
		Params: model.MessageParams{
			Group_id: GroupId,
			Message:  msg,
		},
		Echo: "send_msg",
	}
	jsonData, err := json.Marshal(ws_msg)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	// 通过 WebSocket 发送消息到 OneBot
	err = c.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		fmt.Println("Write Error:", err)
		return err
	}
	return nil
}

func SendGroupMsgEmoji(c *websocket.Conn, Message_id int64) (err error) {
	ws_msg := model.Message{
		Action: "set_msg_emoji_like",
		Params: model.SetMsgEmoji{
			Message_id: int32(Message_id),
			Emoji_id:   1,
			Set:        true,
		},
		Echo: "set_msg_emoji_like",
	}
	jsonData, err := json.Marshal(ws_msg)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	// 通过 WebSocket 发送消息到 OneBot
	err = c.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		fmt.Println("Write Error:", err)
		return err
	}
	return nil
}
