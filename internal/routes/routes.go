package routes

import (
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.homeHandler)
	http.HandleFunc("/", handlers.errorHandler)

	fs := http.FileServer(http.Dir("web/static"))

	http.Handle("/static/", http.StripPrefix("/static", fs))

}
