package repository

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3" // Init the SQLite driver
)

type sqliteRepository struct {
	db *sql.DB
}

func NewSqliteRepository() Repository {

	db, err := sql.Open("sqlite3", "./url_shortener.db")
	if err != nil {
		slog.Error("failed to open database: ", err)
		return nil
	}

	// Create a table
	createTableSQL := `CREATE TABLE IF NOT EXISTS url (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "long_url" TEXT,
        "short_url" TEST
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		slog.Error("Create table error: ", err)
		return nil
	}

	return &sqliteRepository{
		db: db,
	}
}

func (r *sqliteRepository) Close() {
	r.db.Close()
}

func (r *sqliteRepository) Save(shortURL string, longURL string) error {
	//TODO handle update
	insertUserSQL := `INSERT INTO url (long_url, short_url) VALUES (?, ?)`
	_, err := r.db.Exec(insertUserSQL, longURL, shortURL)
	if err != nil {
		return err
	}
	return nil
}

func (r *sqliteRepository) Load(shortURL string) (string, bool, error) {
	query := `SELECT long_url FROM url WHERE short_url = ?`
	row := r.db.QueryRow(query, shortURL)

	var longURL string
	err := row.Scan(&longURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", false, nil
		} else {
			slog.Error("Error when reading from db: ", err)
		}
		return "", false, err
	}

	return longURL, true, nil
}
