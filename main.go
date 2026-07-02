package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var version = "dev"

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "ok")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, version)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file") // source

	// sanitizer: 파일명만 추출해 상위 경로 탈출을 차단
	clean := filepath.Base(filepath.Clean(filename))
	safePath := filepath.Join("./downloads", clean)

	// 허용 디렉터리를 벗어나지 않는지 최종 확인
	if !strings.HasPrefix(safePath, "downloads") {
		http.Error(w, "invalid file", http.StatusBadRequest)
		return
	}

	data, err := os.ReadFile(safePath) // sink (이제 정제된 값만 도달)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	_, _ = w.Write(data)
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/download", downloadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
