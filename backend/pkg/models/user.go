package models

// defines  User data type
type User struct {
	ID          string
	Email       string
	FirstName   string
	LastName    string
	Password    string
	Nickname    string
	About       string
	DateOfBirth string
}

// Repository represent all possible actions availible to deal with User
// all db packages(in case of different db) should implement those function
type UserRepository interface {
	// example code
	Add(User) error // TEST
}
