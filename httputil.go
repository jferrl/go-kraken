package kraken

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
)

func createFakeServer(statusCode int, res string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		http.ServeFile(w, r, filepath.Join("testdata", res))
	}))
}
