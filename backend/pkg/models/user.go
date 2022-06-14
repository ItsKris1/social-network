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
	ImagePath   string
}

// Repository represent all possible actions availible to deal with User
// all db packages(in case of different db) should implement those function
type UserRepository interface {
	Add(User) error //save new user in db
	EmailNotTaken(email string) (bool, error) //returns true if not taken
}
