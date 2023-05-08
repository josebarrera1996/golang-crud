package main

// paquetes
import (
	"fmt"
	"golang-crud/config"
	"golang-crud/controllers"
	"golang-crud/helper"
	"golang-crud/repositories"
	"golang-crud/router"
	"golang-crud/services"
	"net/http"
)

// Función principal
func main() {
	fmt.Printf("Server started...")

	// Implementando la conexión con la B.D
	db := config.DatabaseConnection()

	// Implementando el repositorio
	bookRepository := repositories.NewBookRepository(db)

	// Implementando el servicio
	bookService := services.NewBookServiceImpl(bookRepository)

	// Implementando el controlador
	bookController := controllers.NewBookController(bookService)

	// Definiendo el enrutador
	routes := router.NewRouter(bookController)

	// Levantando el servidor
	server := http.Server{Addr: "localhost:8000", Handler: routes}

	// Manejando los errores
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
