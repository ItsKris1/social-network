package models

// defines  User data type
type User struct {
	Name string `json:"name"`
	ID   string `json:"user_id"`
}

// Repository represent all possible actions availible to deal with User
// all db packages(in case of different db) should implement those function
type UserRepository interface {
	// example code
	Add(User) error // TEST
}
