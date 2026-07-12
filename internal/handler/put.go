package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func Put(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "POST method required",
		})
		return
	}

	err := r.ParseMultipartForm(200 << 20) // 200 MB

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "file not found", http.StatusBadRequest)
		return
	}

	defer file.Close()

	os.MkdirAll("temp", 0755)

	dstPath := filepath.Join("temp", header.Filename)

	dst, err := os.Create(dstPath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// resp := map[string]any{
	// 	"success":  true,
	// 	"filename": header.Filename,
	// 	"size":     header.Size,
	// 	"temp":     dstPath,
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(resp)

	cmd := exec.Command("python", "engine/put.py", dstPath)

	output, err := cmd.Output()

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
