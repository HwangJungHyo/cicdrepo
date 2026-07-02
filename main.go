package main

import (
	"fmt"
	"log"
	"net/http"
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
	log.Fatal(http.ListenAndServe(":8080", nil))  // ② error를 실제 처리
}
