package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	converted, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "Conversion failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	filename := fmt.Sprintf("uploads/%s.txt", time.Now().UTC().Format("20060102-150405"))
	err = os.WriteFile(filename, []byte(converted), 0644)
	if err != nil {
		http.Error(w, "Unable to write file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Converted text saved to %s", filename)
}