package main

import (
	"fmt"
	"net/http"
)

var version = "dev" // 나중에 CI에서 ldflags로 주입할 자리

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, version)
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/version", versionHandler)
	http.ListenAndServe(":8080", nil)
}

