package models

type Group struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`

	Member bool `json:"member"` // trure if current user is a member
	Administrator bool `json:"admin"` // true if current user is admin
}

type GroupRepository interface{
	GetAllAndRelations(userId string) ([]Group, error)
}