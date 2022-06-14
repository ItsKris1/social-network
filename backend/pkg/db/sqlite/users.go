package db

import (
	"database/sql"
	"social-network/pkg/models"
)

type UserRepository struct {
	DB *sql.DB
}

/* --------------- functions that interact with with database --------------- */

// Insert new user in db
func (repo *UserRepository) Add(user models.User) error {
	// example code
	stmt, err := repo.DB.Prepare("INSERT INTO users(user_id, email,first_name, last_name, nickname, about, password, birthday, image) values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(user.ID, user.Email, user.FirstName, user.LastName, user.Nickname, user.About, user.Password, user.DateOfBirth, user.ImagePath); err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) EmailNotTaken(email string) (bool, error) {
	row := repo.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ? LIMIT 1", email)
	var result int
	if err := row.Scan(&result); err != nil {
		if err == sql.ErrNoRows {
			return false, err
		}
	}
	if result != 0 {
		return false, nil
	}
	return true, nil
}
