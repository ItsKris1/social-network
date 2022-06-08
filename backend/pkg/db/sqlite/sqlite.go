package db

import (
	"database/sql"
	"log"
	"social-network/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

// initialize db
func InitDB() *sql.DB {
	// code example
	db, err := sql.Open("sqlite3", "./SNDB.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
	// return nil // delete this line whern return db conn
}

// InitRepositories should be called in server.go
// creates connection to db for all rep
func InitRepositories(db *sql.DB) *models.Repositories {
	return &models.Repositories{
		UserRepo:    &UserRepository{DB: db},
		SessionRepo: &SessionRepository{DB: db},
	}
}
