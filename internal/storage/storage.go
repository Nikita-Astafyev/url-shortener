package storage

type URLStorage interface {
	CreateURL(originalURL string) (string, error)
	GetOriginalURL(shortURL string) (string, error)
	IncrementVisits(shortURL string) error
	GetStats(shortURL string) (int, error)
}
