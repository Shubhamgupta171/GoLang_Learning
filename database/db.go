package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var DB *sqlx.DB

func Connect() {
	var err error

	DB, err = sqlx.Open("sqlite", "./app.db")
	if err != nil {
		log.Fatal("Failed to connect to SQLite:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("SQLite ping failed:", err)
	}

	log.Println("SQLite connected")

	createTables()
}

func createTables() {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	`

	DB.MustExec(schema)
	log.Println("🛠️ Users table created or already exists")
}
