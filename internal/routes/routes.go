package routes

import(
	"net/http"
	"github.com/suppiden/OpenWebinars-go/internal/handlers"
)

func RegisterRoutes(){
	http.HandleFunc("/", handlers)
	http.HandleFunc("/error", handlers.homeHandler)

}

