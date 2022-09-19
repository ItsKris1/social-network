package db

import (
	"database/sql"
	"log"
	"social-network/pkg/models"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

// initialize db
func InitDB() *sql.DB {

	db, err := sql.Open("sqlite3", "./tempDB.db")
	// code example
	if err != nil {
		log.Fatal(err)
	}

	err = Migrations(db)
	if err != nil {
		log.Fatal(err)
	}

	// createTables(db) // temporary code
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

func Migrations(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "../backend/pkg/db/migration/sqlite",
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return err
	}

	fmt.Printf("Applied %d migrations to database.db!", n)
	fmt.Println()

	return nil
}