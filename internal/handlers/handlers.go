package handlers

import (
	"html/template"
	"net/http"
	"os"

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
		Title: "OpenWebinars",
		Author: "Jonathan",
		Welcome: "Mira lo que acabo de hacer",
		

	}

	page :=r.URL.Path[1:]
	if page == ""{
		page = "index.html"
	}


	tmplFile := "web/templates/" + page

	if _,err := os.Stat(tmplFile); err !=nil{
		tmplFile = "web/templates/error.html"

		data.ErrorCode = http.StatusNotFound
		data.ErrorMessage = "Pagina no encontrada"
	}

	renderTemplate(w, tmplFile, data)


}


func errorHandler(w http.ResponseWriter, r *http.Request){
	data := models.PAgeData{
		Title: "Pagina no encontrada",
		ErrorMessage: "¡Error interno del servidos!",
		ErrorCode: http.StatusInternalServerError,
	}

	renderTemplate(w, "web/templates/error.html", data)
}