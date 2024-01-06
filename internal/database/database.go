package database

import (
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sqlx.DB {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    // home directory of project
    homeDir,err := os.UserHomeDir()
    dbName := os.Getenv("DATABASE_NAME")

    dbPath := filepath.Join(homeDir, dbName)
    db, err := sqlx.Connect("sqlite3", dbPath)
    if err != nil {
        panic(err)
    }

    schema := `
    CREATE TABLE IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS articles (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        category_id INTEGER,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        FOREIGN KEY (category_id) REFERENCES categories (id)
    );
    `
    
    // Execute the schema or migration SQL
    _, err = db.Exec(schema)
    if err != nil {
        panic(err)
    }

    return db
}