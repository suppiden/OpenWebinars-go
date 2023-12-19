package main
import(
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/suppiden/OpenWebinars-go/internal/config"
	"github.com/suppiden/OpenWebinars-go/internal/routes"

)

func main(){
	logFile, err := os.OpenFile("eror.log", os.O_RDWR|os.O_CREATE| os.O_APPEND, 0666)

	if err !=nil{
		fmt.Println("Error al abrir el archivo de registro: ", err)
		return
	}

	defer logFile.Close()
	log.SetOutput(logFile)



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