package models

// defines  User data type
type User struct {
	ID          string `json:"id"`
	Email       string `json:"login,omitempty"`
	FirstName   string `json:"first-name,omitempty"`
	LastName    string `json:"last-name,omitempty"`
	Password    string `json:"password,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	About       string `json:"about,omitempty"`
	DateOfBirth string `json:"date-of-birth,omitempty"`
	ImagePath   string `json:"avatar,omitempty"`

	Follower  bool `json:"follower,omitempty"`  //if this user is following another user
	Following bool `json:"following,omitempty"` //if curr user is following this one
}

// Repository represent all possible actions availible to deal with User
// all db packages(in case of different db) should implement those function
type UserRepository interface {
	Add(User) error                           //save new user in db
	EmailNotTaken(email string) (bool, error) //returns true if not taken
	FindUserByEmail(email string) (User, error)
	GetAllAndFollowing(userID string) ([]User, error) //all users and follow info
	GetDataMin(userID string) (User, error)       // returns id, nickname and image
}
