package ws

import "social-network/pkg/models"

// represent websocket server
type Server struct {
	Clients map[*Client]bool
	Repos   *models.Repositories
}

func StartServer(repos *models.Repositories) *Server {
	server := &Server{
		make(map[*Client]bool),
		repos,
	}
	return server
}

// register client
func (s *Server) RegisterNewClient(client *Client) {
	s.Clients[client] = true //update client list

	messageToNewClient := WsMessage{
		Action:  OnlineUsersAction,
		OnlineUsers: s.GetOnlineUsers(),
	}
	client.send <- messageToNewClient.encode()
	
	message := WsMessage{
		UserID: client.ID,
		Action:       UserJoinAction,
		Message: client.ID,
	}
	s.broadcast(message)

}

// register and unregister clients
func (s *Server) UnregisterClient(client *Client) {
	if _, ok := s.Clients[client]; ok {
		delete(s.Clients, client)
	}
	message := WsMessage{
		UserID: client.ID,
		Action:       UserLeftAction,
		Message: client.ID,
	}
	s.broadcast(message)
}

func (server *Server) broadcast(message WsMessage){
	for client := range server.Clients{
		client.send <- message.encode()
	}
}
func (s *Server) GetOnlineUsers() []string{
	var onlineUsers = []string{}
	for client := range s.Clients{
		onlineUsers = append(onlineUsers, client.ID)
	}
	return onlineUsers
}