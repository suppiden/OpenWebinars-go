package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenido a la pagina de inicio")
}
func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aquí se muestran los productos")
}
func productDetailtHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Path[len("/product/detail/"):]
	fmt.Fprintf(w, "Detalles del producto con ID: %s", productID)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/product", productHandler)
	http.HandleFunc("/product/detail/", productDetailtHandler)

	fmt.Println("Servidor está corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
