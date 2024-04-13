package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type PageData struct {
	Title        string
	Message      template.HTML
	ErrorCode    int
	ErrorMessage string
}

func renderTemplate(w http.ResponseWriter, tmplFile string, data PageData) {
	tmpl, err := template.ParseFiles(tmplFile)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Diego Alonso Molina",
		Message: template.HTML(`
		La mejor forma de hacer una página web es con <b>"Go"</b>.
		`),
	}

	tmplFile := filepath.Join("web/templates/", "home.html")
	renderTemplate(w, tmplFile, data)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:        "Error 404",
		ErrorCode:    404,
		ErrorMessage: "¡Página no encontrada!",
	}

	tmplFile := filepath.Join("web/templates/", "error.html")
	renderTemplate(w, tmplFile, data)
}

func main() {

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/error", errorHandler)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
