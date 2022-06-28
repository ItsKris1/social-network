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
	// return nil // delete this line where return db conn
}

// InitRepositories should be called in server.go
// creates connection to db for all rep
func InitRepositories(db *sql.DB) *models.Repositories {
	return &models.Repositories{
		UserRepo:    &UserRepository{DB: db},
		SessionRepo: &SessionRepository{DB: db},
		GroupRepo:   &GroupRepository{DB: db},
		PostRepo:    &PostRepository{DB: db},
		CommentRepo: &CommentRepository{DB: db},
		NotifRepo:   &NotifRepository{DB: db},
	}
}

func createTables(db *sql.DB) {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS users (
			"user_id" varchar(255) not null,
			"created_at" datetime not null default CURRENT_TIMESTAMP,
			"email" varchar(255) not null,
			"first_name" varchar(255) not null,
			"last_name" varchar(255) not null,
			"nickname" varchar(255) null,
			"birthday" datetime not null,
			"image" varchar(255) null,
			"about" TEXT null,
			"status" varchar(255) not null default PUBLIC,
			"password" varchar(255) not null,
			primary key ("user_id")
		)`,
		`
		 CREATE TABLE IF NOT EXISTS groups (
			"group_id" VARCHAR(255) not null,
			"administrator" VARCHAR(255) not null,
			"name" VARCHAR(255) not null,
			"description" VARCHAR(255) null,
			primary key ("group_id")
		)`,
		`
		CREATE TABLE IF NOT EXISTS posts (
			"post_id" VARCHAR(255) not null,
			"group_id" varchar(255) null,
			"created_by" varchar(255) not null,
			"created_at" datetime not null default CURRENT_TIMESTAMP,
			"content" TEXT null,
			"image" varchar(255) null,
			"visibility" varchar(255) null default PUBLIC,
			primary key ("post_id")
 		)`,
		`
		CREATE TABLE IF NOT EXISTS comments (
			"comment_id" VARCHAR(255) not null,
			"post_id" VARCHAR(255) not null,
			"created_at" datetime not null default CURRENT_TIMESTAMP,
			"created_by" varchar(255) not null,
			"content" text null,
			"image" varchar(255) null,
			primary key ("comment_id")
		)`,
		`
		CREATE TABLE IF NOT EXISTS event (
			"event_id" VARCHAR(255) not null,
			"group_id" VARCHAR(255) not null,
			"created_by" VARCHAR(255) not null,
			"created_at" DATETIME not null default CURRENT_TIMESTAMP,
			"title" VARCHAR(255) not null,
			"content" VARCHAR(255) not null,
			"date" DATETIME not null,
			primary key ("event_id")
		)`,
		`
		CREATE TABLE IF NOT EXISTS messages (
			"message_id" VARCHAR(255) not null,
			"sender_id" VARCHAR(255) not null,
			"receiver_id" DATETIME not null default CURRENT_TIMESTAMP,
			"type" VARCHAR(255) not null,
			"created_at" VARCHAR(255) not null,
			"content" VARCHAR(255) not null,
			primary key ("message_id")
		)`,
		`CREATE TABLE IF NOT EXISTS sessions (
			"session_id" VARCHAR(255) NOT NULL PRIMARY KEY,
			"user_id" VARCHAR(255) NOT NULL,
			"expiration_time" DATETIME NOT NULL
		)`,
		`
		CREATE TABLE IF NOT EXISTS almost_private (
			"user_id" VARCHAR(255) not null,
			"post_id" VARCHAR(255) not null
		)`,
		`
		CREATE TABLE IF NOT EXISTS event_users (
			"event_id" VARCHAR(255) not null,
			"user_id" VARCHAR(255) not null,
		)`,
		`
		CREATE TABLE IF NOT EXISTS group_users (
			"group_id" VARCHAR(255) not null,
			"user_id" VARCHAR(255) not null,
		)`,
		`
		CREATE TABLE IF NOT EXISTS followers (
			"follower_id" VARCHAR(255) not null,
			"user_id" VARCHAR(255) not null,
		)`,
	}
	for _, table := range tables {
		statement, _ := db.Prepare(table)
		defer statement.Close()
		statement.Exec()
	}
}
