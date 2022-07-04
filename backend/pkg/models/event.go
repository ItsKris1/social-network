package models

type Event struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Date     string `json:"date"`
	GroupID  string `json:"groupId"`
	AuthorID string `json:"authorId"`
}

type EventRepository interface {
	GetAll(groupId string) ([]Event, error) //get all events for group
}
