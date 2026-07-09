package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok - env=%s\n", os.Getenv("ENV_NAME"))
	})
	http.ListenAndServe(":8080", nil)
}
// trigger build
// qa promote trigger
