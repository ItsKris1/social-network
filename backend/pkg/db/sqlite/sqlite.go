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
		GroupRepo:   &GroupRepository{DB: db},
		// PostRepo: &PostRepository{DB: db}
	}
}

func createTables(db *sql.DB) {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS users (
			"user_id" VARCHAR(255) not null,
			"created_at" datetime not null default CURRENT_TIMESTAMP,
			"email" varchar(255) not null,
			"first_name" varchar(255) not null,
			"last_name" varchar(255) not null,
			"nickname" varchar(255) null,
			"image" varchar(255) null,
			"about" TEXT null,
			"status" varchar(255) not null default PUBLIC,
			"password" varchar(255) not null,
			"birthday" DATETIME not null,
			primary key ("user_id")
		)`,
		`CREATE TABLE IF NOT EXISTS sessions (
			"user_id" VARCHAR(255) NOT NULL,
			"session_id" VARCHAR(255) NOT NULL PRIMARY KEY,
			"expiration_time" DATETIME NOT NULL
		)`,
		`
		CREATE TABLE IF NOT EXISTS posts (
			"post_id" VARCHAR(255) not null,
			"created_at" datetime not null default CURRENT_TIMESTAMP,
			"group_id" varchar(255) null,
			"created_by" varchar(255) not null,
			"content" TEXT null,
			"image" varchar(255) null,
			"visibilitY" varchar(255) null default PUBLIC,
			primary key ("post_id")
 		)`,
		`
		 CREATE TABLE IF NOT EXISTS groups (
			"group_id" VARCHAR(255) not null,
			"name" VARCHAR(255) not null,
			"description" VARCHAR(255) null,
			"administrator" VARCHAR(255) not null,
			primary key ("group_id")
		)`,
		`
		CREATE TABLE IF NOT EXISTS group_users (
			"group_id" VARCHAR(255) not null,
			"user_id" VARCHAR(255) not null
		)`,
		`
		CREATE TABLE IF NOT EXISTS followers (
			"user_id" VARCHAR(255) not null,
			"follower_id" VARCHAR(255) not null
		)`,
		`
		CREATE TABLE IF NOT EXISTS comments (
			"comment_id" VARCHAR(255) not null,
			"created_at" datetime not null default CURRENT_TIMESTAMP,
			"created_by" varchar(255) not null,
			"content" text null,
			"image" varchar(255) null,
			primary key ("comment_id")
		)`,
	}
	for _, table := range tables {
		statement, _ := db.Prepare(table)
		defer statement.Close()
		statement.Exec()
	}
}
