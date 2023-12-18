package routes

import(
	"net/http"
	"github.com/suppiden/OpenWebinars-go/internal/handlers"
)

func RegisterRoutes(){
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/error", handlers.errorHandler)

}

