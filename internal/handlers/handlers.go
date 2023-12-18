package handlers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/suppiden/OpenWebinars-go/internal"
	"github.com/suppiden/OpenWebinars-go/internal/models"
)




func renderTemplate(w http.ResponseWriter, tmplFile string, data interface{} ){
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
	data := models.PAgeData{
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