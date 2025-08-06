package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type MockStorage struct {
	CreateURLError    error
	GetOriginalURLVal string
	GetOriginalURLErr error
}

func (m *MockStorage) CreateURL(originalURL string) (string, error) {
	return "abc123", m.CreateURLError
}

func (m *MockStorage) GetOriginalURL(shortURL string) (string, error) {
	if m.GetOriginalURLErr != nil {
		return "", m.GetOriginalURLErr
	}
	return m.GetOriginalURLVal, nil
}

func (m *MockStorage) IncrementVisits(shortURL string) error {
	return nil
}

func (m *MockStorage) GetStats(shortURL string) (int, error) {
	return 0, nil
}

func TestCreateShortURL_Success(t *testing.T) {
	storage := &MockStorage{}
	handler := NewURLHandler(storage)

	req := httptest.NewRequest("POST", "/create", strings.NewReader("url=https://google.com"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	handler.CreateShortURL(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "abc123", w.Body.String())
}
