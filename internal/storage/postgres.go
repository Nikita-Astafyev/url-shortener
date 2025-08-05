package storage

import (
	"database/sql"
	"fmt"

	"github.com/Nikita-Astafyev/url-shortener/internal/config"
	"github.com/Nikita-Astafyev/url-shortener/internal/service"
)

type PostgresStorage struct {
	DB *sql.DB
}

func NewPostgresStorage(cfg config.DBConfig) (*PostgresStorage, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{DB: db}, nil
}

func (s *PostgresStorage) Init() error {
	_, err := s.DB.Exec(`
        CREATE TABLE IF NOT EXISTS urls (
            id SERIAL PRIMARY KEY,
            original_url TEXT NOT NULL,
            short_url TEXT NOT NULL UNIQUE,
            visits INTEGER DEFAULT 0,
            created_at TIMESTAMP DEFAULT NOW()
        );
    `)
	return err
}

func (s *PostgresStorage) CreateURL(originalURL string) (string, error) {
	shortURL := service.GenerateShortURL()
	_, err := s.DB.Exec(
		"INSERT INTO urls (original_url, short_url) VALUES ($1, $2)",
		originalURL, shortURL,
	)

	return shortURL, err
}

func (s *PostgresStorage) GetOriginalURL(shortURL string) (string, error) {
	var originalURL string
	err := s.DB.QueryRow(
		"SELECT original_url FROM urls WHERE short_url = $1",
		shortURL,
	).Scan(&originalURL)

	if err == sql.ErrNoRows {
		return "", fmt.Errorf("URL not found")
	}
	if err != nil {
		return "", fmt.Errorf("DB error: %v", err)
	}

	return originalURL, nil
}

func (s *PostgresStorage) IncrementVisits(shortURL string) error {
	_, err := s.DB.Exec(
		"UPDATE urls SET visits = visits + 1 WHERE short_url = $1",
		shortURL,
	)
	return err
}

func (s *PostgresStorage) GetStats(shortURL string) (int, error) {
	var visits int
	err := s.DB.QueryRow(
		"SELECT visits FROM urls WHERE short_url = $1",
		shortURL,
	).Scan(&visits)
	return visits, err
}
