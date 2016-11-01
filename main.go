package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/nanopx/go-glide-docker/handlers"
	"fmt"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.Handle("/status", handlers.NotImplemented).Methods("GET")

	fmt.Println("Serving...")
	http.ListenAndServe(":3000", r)
}
