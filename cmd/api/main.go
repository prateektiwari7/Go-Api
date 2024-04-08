package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/prateektiwari7/Go-Api/internal/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
}
