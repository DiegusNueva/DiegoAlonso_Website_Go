package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("web/templates/", "home.html"))
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("web/templates/", "error.html"))
	})

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
