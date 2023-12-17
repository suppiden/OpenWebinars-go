package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"html/template"
)

type PageData struct{
	Title string
	Message template.HTML
	ErrorCode int
	ErrorMessage string
}

func renderTemplate(w http.ResponseWriter, tmplFile string, data PageData ){
	tmpl, err := template.ParseFiles(tmplFile)

	if err !=nil{
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

		err = tmpl.Execute(w, data)


	if err !=nil{
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request){
	data := PageData{
		Title: "OpenWebinars / Jony",
		
		Message: template.HTML("La plataforma donde encontraras el curso <b>De Go hecha por mi<b>"),

	}

	tmplFile := filepath.Join("web/templates/", "index.html")
	renderTemplate(w, tmplFile, data)
}


func errorHandler(w http.ResponseWriter, r *http.Request){
	data := PageData{
		Title: "error 404",
		ErrorMessage: "¡Pagina no encontrada!",
		ErrorCode: 404,
	}

	tmplFile := filepath.Join("web/templates/", "error.html")
	renderTemplate(w, tmplFile, data)
}
func main() {
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))


	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/error", errorHandler)




	fmt.Println("Servidor está corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
