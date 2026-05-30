package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type response struct {
	Application string `json:"application"`
	Task        string `json:"task"`
	Status      string `json:"status"`
	Time        string `json:"time"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(response{
			Application: "Zad2",
			Task:        "GitHub Actions, GHCR, DockerHub cache, CVE test",
			Status:      "OK",
			Time:        time.Now().UTC().Format(time.RFC3339),
		})
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK\n"))
	})

	addr := ":" + port
	log.Printf("Serwer uruchomiony na porcie %s", port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
