package handlers

import (
	"net/http"
	"os/filepath"
	"strings"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(".", "static", r.URL.Path[1:])

	if !strings.HasPrefix(filepath.Clean(path), filepath.Clean(filepath.Join(".", "static"))) {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}

	http.ServeFile(w, r, path)
}
