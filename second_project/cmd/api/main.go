package main

import (
	"fmt"
	"net/http"

	"second_project/internal/handler"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()

	handler.Handler(r)

	fmt.Println("Starting GO API...")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Error(err)
	}

	fmt.Printf("GO API started on port 8080\n")
}
