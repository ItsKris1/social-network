package models

type ChatMessage struct {
	ID         string `json:"id"`
	SenderId   string `json:"senderId"`
	ReceiverId string `json:"receiverId"`
	Type       string `json:"type"` //GROUP|PERSON
	Content    string `json:"content"`

	Sender User `json:"sender"`
}

type MsgRepository interface {
	Save(ChatMessage) error
	//get all for specific chat
	// needs  RECEIVER and SENDER as input
	GetAll(ChatMessage) ([]ChatMessage, error)
	GetAllGroup(userId, groupId string) ([]ChatMessage, error)
}
