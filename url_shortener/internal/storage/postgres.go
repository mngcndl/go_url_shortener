package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %v", err)
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Save(short, original string) error {
	_, err := s.db.Exec("INSERT INTO urls (short, original) VALUES ($1, $2)", short, original)
	return err
}

func (s *PostgresStorage) Get(short string) (string, bool, error) {
	var original string
	err := s.db.QueryRow("SELECT original FROM urls WHERE short = $1", short).Scan(&original)
	if err == sql.ErrNoRows {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return original, true, nil
}