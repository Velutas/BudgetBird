package main

import (
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()


	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a webpage"))
		log.Info("Page accessed")
	})


	http.ListenAndServe(":8080", r)
}