package models

type Notification struct {
	ID       string `json:"id"`
	TargetID string
	Type     string
	Content  string
}

type NotifRepository interface {
	Save(Notification) error
	Delete(notificationId string) error
	CheckIfExists(Notification)(bool, error) // true if exists, false otherwise

	//get all pending requests to join group
	GetGroupRequestsIds(groupId string) ([]string, error) 
	// get specific user_id from request to join
	GetUserFromRequest(notificationId string)(string, error)
	// get group id from specific request
	GetGroupId(notificationId string)(string, error)
}
