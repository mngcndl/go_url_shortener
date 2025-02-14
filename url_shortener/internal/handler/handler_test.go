package handler

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
)

type mockService struct{}

func (m *mockService) CreateShortURL(original string) (string, error) {
    return "abcd1234", nil
}

func (m *mockService) GetOriginalURL(short string) (string, error) {
    return "http://example.com/some/long/url", nil
}

func TestCreateShortURL(t *testing.T) {
    service := &mockService{}
    h := NewHandler(service)

    req, err := http.NewRequest("POST", "/shorten", bytes.NewBufferString("url=http://example.com/some/long/url"))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    recorder := httptest.NewRecorder()
    h.CreateShortURL(recorder, req)

    assert.Equal(t, http.StatusOK, recorder.Code)

    var response map[string]string
    err = json.Unmarshal(recorder.Body.Bytes(), &response)
    assert.NoError(t, err)
    assert.Equal(t, "abcd1234", response["short"])
    assert.Equal(t, "http://example.com/some/long/url", response["original"])
}

func TestGetOriginalURL(t *testing.T) {
    service := &mockService{}
    h := NewHandler(service)

    req, err := http.NewRequest("GET", "/abcd1234", nil)
    if err != nil {
        t.Fatal(err)
    }

    recorder := httptest.NewRecorder()
    h.GetOriginalURL(recorder, req)

    assert.Equal(t, http.StatusFound, recorder.Code)
    assert.Equal(t, "http://example.com/some/long/url", recorder.Header().Get("Location"))
}
