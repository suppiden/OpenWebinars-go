package main
import(
	"fmt"
	"log"
	"net/http"

	"github.com/suppiden/OpenWebinars-go/internal/config"
	"github.com/suppiden/OpenWebinars-go/internal/routes"

)

func main(){
	cfg := config.LoadConfig()

	routes.RegisterRoutes()

	addr := fmt.Sprintf(":%s", cfg.Port)

	server := &http.Server{
		Addr: addr,
		Handler: nil,
	}

	fmt.Println("Servidor corriendo en localhost:" + cfg.Port)
	log.Fatal(server.ListenAndServe())
}