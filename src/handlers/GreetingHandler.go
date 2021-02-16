package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mattias-fjellstrom/go-custom-handler-logging/src/logging"
)

// GreetingHandler handles GET requests to /api/greeting
func GreetingHandler(w http.ResponseWriter, r *http.Request) {
  logger := logging.NewLogger()

  logger.Info("Log before a five seconds sleep")
  time.Sleep(5 * time.Second)
  logger.Info("Log after a five seconds sleep")
  logger.Warning("This is a warning log")
  logger.Error("This is an error log")

  response := "Hello!"
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(response)
}