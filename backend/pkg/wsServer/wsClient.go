package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"social-network/pkg/models"
	"social-network/pkg/utils"

	"github.com/gorilla/websocket"
)

// represents single websocket client
type Client struct {
	ID       string
	conn     *websocket.Conn
	wsServer *Server
	send     chan []byte
}

func NewClient(conn *websocket.Conn, wsServer *Server, ID string) *Client {
	return &Client{
		ID:       ID,
		conn:     conn,
		wsServer: wsServer,
		send:     make(chan []byte, 256),
	}
}

/* ----------- main enterypoint for dealing with incoming messages ---------- */
func (client *Client) handleNewMessage(msgJSON []byte) {
	var msg WsMessage
	if err := json.Unmarshal(msgJSON, &msg); err != nil {
		log.Println(err)
		return
	}

	switch msg.Action {
	case DisconnectAction:
		client.disconnect()
	}
}

/* -------------------------------------------------------------------------- */
/*                           client action functions                          */
/* -------------------------------------------------------------------------- */
func (client *Client) disconnect() {
	client.wsServer.UnregisterClient(client)
}

func (client *Client) SendNotification(notif models.Notification) {
	switch notif.Type {
	case "GROUP_INVITE":
		notif.Group, _ = client.wsServer.repos.GroupRepo.GetData(notif.Content)
		notif.User, _ = client.wsServer.repos.UserRepo.GetDataMin(notif.Sender)
	case "FOLLOW":
		notif.User, _ = client.wsServer.repos.UserRepo.GetDataMin(notif.Content)
	case "EVENT":
		notif.Event, _ = client.wsServer.repos.EventRepo.GetData(notif.Content)
		notif.User, _ = client.wsServer.repos.UserRepo.GetDataMin(notif.Sender)
	case "GROUP_REQUEST":
		notif.User, _ = client.wsServer.repos.UserRepo.GetDataMin(notif.Content)
		notif.Group, _ = client.wsServer.repos.GroupRepo.GetData(notif.TargetID)
	}
	/* ---------------------------- add message text ---------------------------- */
	utils.DefineNotificationMsg(&notif)

	/* ---------------------------- construct message --------------------------- */
	message := WsMessage{
		Action:       NotificationAction,
		Notification: notif,
	}
	/* ---------------------------------- send ---------------------------------- */
	client.send <- message.encode()
}

func (client *Client) SendChatMessage(msg models.ChatMessage) {
	message := WsMessage{
		Action:      ChatAction,
		ChatMessage: msg,
	}
	client.send <- message.encode()
}

/* -------------------------------------------------------------------------- */
/*                 basic reader and writer for websocket conn                 */
/* -------------------------------------------------------------------------- */
// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func (client *Client) Reader() {
	for {
		// read in a message
		_, msg, err := client.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		client.handleNewMessage(msg)
	}
}

// define a writer which will send
// new messages to our WebSocket
// endpoint
func (client *Client) Writer() {
	for {
		message, ok := <-client.send
		if !ok {
			log.Println("err on writing message")
			return
		}
		if err := client.conn.WriteJSON(string(message)); err != nil {
			fmt.Println(err)
			return
		}
	}
}
