package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"service"
	"strings"
	"time"
)

func handleMain(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "method not allowed", http.StatusInternalServerError)
		return
	}
	http.ServeFile(res, req, "index.html")
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusInternalServerError)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10MB
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "bad form: "+err.Error(), http.StatusInternalServerError)
		return
	}

	f, hdr, err := r.FormFile("myFile")

	if err != nil {
		http.Error(w, "file missing: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer f.Close()

	data, err := io.ReadAll(f)

	if err != nil {
		http.Error(w, "convert error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	convert, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "convert error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	base := filepath.Base(hdr.Filename)
	ext := strings.ToLower(filepath.Ext(base))

	ts := time.Now().UTC().Format(time.RFC3339Nano)

	file_name := fmt.Sprintf("out_%s%s", ts, ext)

	if err := os.WriteFile(file_name, []byte(convert), 0644); err != nil {
		http.Error(w, "convert error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(convert))
}

func NewRouter() *http.ServeMux {
	route := http.NewServeMux()
	route.HandleFunc("/", handleMain)
	route.HandleFunc("/upload", handleUpload)
	return route
}
