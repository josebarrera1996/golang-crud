package router

import (
	"fmt"
	"golang-crud/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Rutas
func NewRouter(bookController *controllers.BookController) *httprouter.Router {

	// Definiendo el enrutador
	router := httprouter.New()

	/* Definiendo las rutas */
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})
	router.GET("/api/books", bookController.FindAll)
	router.GET("/api/books/:bookId", bookController.FindById)
	router.POST("/api/books", bookController.Create)
	router.PATCH("/api/books/:bookId", bookController.Update)
	router.DELETE("/api/books/:bookId", bookController.Delete)

	return router
}
