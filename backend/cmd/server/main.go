package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	addr := ":8080"
	logger.Info("server starting", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
