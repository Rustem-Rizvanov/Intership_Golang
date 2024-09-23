package repository

import (
	"database/sql"
)

type PostgresHashRepository struct {
	db *sql.DB
}

func NewPostgresHashRepository(dataSourceName string) (*PostgresHashRepository, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS hash_table (
		id SERIAL PRIMARY KEY,
		word TEXT NOT NULL,
		hash TEXT NOT NULL
	);
	`)
	if err != nil {
		return nil, err
	}

	return &PostgresHashRepository{db: db}, nil
}

func (r *PostgresHashRepository) Save(word string, hash string) error {
	_, err := r.db.Exec("INSERT INTO hash_table (word, hash) VALUES ($1, $2)", word, hash)
	return err
}
