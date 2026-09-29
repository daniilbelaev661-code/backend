package handlers

import (
	"io"
	"net/http"
	"os"
)

// SaveHandler — POST /api/save
// принимает текст и дописывает его в data.txt
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Read error", http.StatusBadRequest)
		return
	}

	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "File error", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(string(body) + "\n"); err != nil {
		http.Error(w, "Write error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// ReadHandler — GET /api/read
// возвращает содержимое data.txt
func ReadHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		if os.IsNotExist(err) {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(""))
			return
		}
		http.Error(w, "Read error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(data)
}