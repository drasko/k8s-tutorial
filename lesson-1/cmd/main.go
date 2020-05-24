package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

var logger *zap.Logger

type Status struct {
	Running bool `json:"running"`
}

func main() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		fmt.Println("Failed to initialize Logger")
		return
	}
	defer logger.Sync()

	r := chi.NewRouter()
	r.Get("/status", status)

	fmt.Println("Starting server at port 3333")
	if err := http.ListenAndServe(":3333", r); err != nil {
		fmt.Println("Error serving HTTP")
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	logger.Info("status called", zap.String("method", r.Method))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	b, err := json.Marshal(Status{Running: true})
	if err != nil {
		fmt.Println("Error marshaling JSON")
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("Error writing response")
		return
	}
}
