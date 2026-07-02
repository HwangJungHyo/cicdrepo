package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var version = "dev"

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "ok")        // ① 의도적 무시를 명시
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, version)     // ① 동일
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/download", downloadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))  // ② error를 실제 처리
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file") // source: 사용자 입력
	data, err := os.ReadFile(filename)    // sink: 파일 시스템 접근
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	_, _ = w.Write(data)
}
