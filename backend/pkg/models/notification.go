package models

type Notification struct {
	ID       string `json:"id"`
	TargetID string
	Type     string
	Content  string
}

type NotifRepository interface {
	Save(Notification) error
}
