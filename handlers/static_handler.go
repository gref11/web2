package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (h *Handler) StaticHandler(w http.ResponseWriter, r *http.Request) {
	// Базовый каталог со статикой
	staticDir := filepath.Join(".", "static")

	// Получаем путь к файлу и нормализуем его
	requestPath := r.URL.Path
	if requestPath == "/" {
		requestPath = "/index.html"
	}

	// Строим полный путь к файлу
	fullPath := filepath.Join(staticDir, requestPath)
	fullPath = filepath.Clean(fullPath)

	// Проверяем, что путь находится внутри staticDir
	if !strings.HasPrefix(fullPath, filepath.Clean(staticDir)+string(filepath.Separator)) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Проверяем существование файла
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	// Отдаём файл
	http.ServeFile(w, r, fullPath)
}
