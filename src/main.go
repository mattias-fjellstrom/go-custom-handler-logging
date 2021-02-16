package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mattias-fjellstrom/go-custom-handler-logging/src/handlers"
)

func main() {
  port, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
  if !exists {
    port = "8080"
  }

  mux := http.NewServeMux()
  mux.HandleFunc("/api/greeting", handlers.GreetingHandler)

  log.Fatal(http.ListenAndServe(":"+port, mux))
}