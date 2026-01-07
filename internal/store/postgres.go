package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Url struct {
	ID        int
	Original  string
	ShortCode string
	CreatedAt time.Time
}

type Store struct {
	db *sql.DB
}

func NewStore(connString string) (*Store, error) {
	db, err := sql.Open("pgx", connString)

	if err != nil {
		return  nil, fmt.Errorf("Erro: falha ao abrir conexão: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Erro: falha ao estabelecer conexão: %w", err)
	}

	s := &Store{db: db}

	if err := s.createTable(); err != nil {
		return nil, fmt.Errorf("Erro: falha ao criar a tabela de links: %w", err)
	}

	return s, nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		original_url TEXT NOT NULL,
		short_code VARCHAR(10) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query)
	return err
}