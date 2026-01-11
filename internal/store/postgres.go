package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Driver do Postgres
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
	// Abre a conexão com o banco de dados PostgreSQL
	db, err := sql.Open("pgx", connString)

	if err != nil {
		return  nil, fmt.Errorf("Erro: falha ao abrir conexão: %w", err)
	}

	// Verifica se a conexão está ativa
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Erro: falha ao estabelecer conexão: %w", err)
	}

	// Cria a instância da Store
	s := &Store{db: db}

	// Cria a tabela de links, se não existir
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

	// Define um contexto com timeout para a execução da query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Executa a query de criação da tabela
	_, err := s.db.ExecContext(ctx, query)
	return err
}

func (s *Store) Get(code string) (string, error) {
	var url string 

	query := `
			SELECT original_url
			FROM urls
			WHERE short_code = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, code).Scan(&url)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Erro: URL não encontrada: %w", err)
		}
		return "", fmt.Errorf("Erro: falha ao buscar pela URL: %w", err)
	}

	return url, nil
}

func (s *Store) Save(originalUrl string, shortCode string) (int, error) {
	query := `
			INSERT INTO urls (original_url, short_code)
			VALUES ($1, $2)
			RETURNING id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id int
	err := s.db.QueryRowContext(ctx, query, originalUrl, shortCode).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("Erro: falha ao salvar url: %w", err)
	}

	return id, nil
}