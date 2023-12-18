package routes

import(
	"net/http"
	"github.com/suppiden/OpenWebinars-go/internal/handlers"
)

func RegisterRoutes(){
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/error", handlers.ErrorHandler)

	fs := http.FileServer((http.Dir("web/static")))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

