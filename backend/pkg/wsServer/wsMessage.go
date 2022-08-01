package ws

import (
	"encoding/json"
	"log"
	"social-network/pkg/models"
)

/* --------------------- actions for websocket messages --------------------- */
const DisconnectAction = "disconnect" //disconnect from websocket server
const NotificationAction = "notification"
const ChatAction = "chat"

type WsMessage struct {
	Action       string              `json:"action"`  //msg request action
	Message      string              `json:"message"` //The actual message
	Notification models.Notification `json:"notification"`
	ChatMessage  models.ChatMessage  `json:"chatMessage"`
}

// encode method that can be called to create a json []byte object
func (message *WsMessage) encode() []byte {
	json, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return json
}
