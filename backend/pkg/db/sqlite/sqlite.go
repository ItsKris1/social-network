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
	db, err := sql.Open("sqlite3", "./tempDB.db")
	if err != nil {
		log.Fatal(err)
	}
	createTables(db) // temporary code
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

func createTables(db *sql.DB) {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS users (
			"user_id" VARCHAR(255) NOT NULL PRIMARY KEY,
			"username" VARCHAR(255) NOT NULL,
			"created_at" DATETIME default CURRENT_TIMESTAMP)`,

		`CREATE TABLE IF NOT EXISTS sessions (
			"user_id" VARCHAR(255) NOT NULL,
			"session_id" VARCHAR(255) NOT NULL PRIMARY KEY,
			"expiration_time" DATETIME NOT NULL
		)`,
	}
	for _, table := range tables {
		statement, _ := db.Prepare(table)
		defer statement.Close()
		statement.Exec()
	}
}
