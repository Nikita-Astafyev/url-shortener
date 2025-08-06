package storage

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *PostgresStorage {
	db, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)

	storage := &PostgresStorage{DB: db}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			id SERIAL PRIMARY KEY,
			original_url TEXT NOT NULL,
			short_url TEXT NOT NULL UNIQUE,
			visits INTEGER DEFAULT 0
		);
	`)
	require.NoError(t, err)

	t.Cleanup(func() {
		db.Close()
	})

	return storage
}

func TestCreateAndGetURL(t *testing.T) {
	storage := setupTestDB(t)

	originalURL := "https://google.com"
	shortURL, err := storage.CreateURL(originalURL)
	require.NoError(t, err)
	require.NotEmpty(t, shortURL)

	foundURL, err := storage.GetOriginalURL(shortURL)
	require.NoError(t, err)
	require.Equal(t, originalURL, foundURL)
}

func TestGetNonExistentURL(t *testing.T) {
	storage := setupTestDB(t)

	_, err := storage.GetOriginalURL("nonexistent")
	require.Error(t, err)
}
