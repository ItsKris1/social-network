package models

type Event struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Date     string `json:"date"`
	GroupID  string `json:"groupId"`
	AuthorID string `json:"authorId"`
	// going holds status value if user going to event or not
	Going string `json:"going"` // YES || NO
}

type EventRepository interface {
	GetAll(groupId string) ([]Event, error)      //get all events for group
	Save(Event) error                            // save new event
	AddParticipant(eventID, userID string) error // save new participant
	RemoveParticipant(eventID, userID string) error // remove participant
	IsParticipating(eventID, userID string) (bool,error)
}
