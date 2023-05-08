package controllers

import (
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Controlador para manejar la interfaz REST de la lógica comercial
// Implemenetando los servicios
type BookController struct {
	BookService services.BookService
}

func NewBookController(bookService services.BookService) *BookController {
	return &BookController{BookService: bookService}
}

/* Implementando los métodos */
func (controller *BookController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookCreateRequest := request.BookCreateRequest{}
	helper.ReadRequestBody(requests, &bookCreateRequest)

	controller.BookService.Create(requests.Context(), bookCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *BookController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookUpdateRequest := request.BookUpdateRequest{}
	helper.ReadRequestBody(requests, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id

	controller.BookService.Update(requests.Context(), bookUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *BookController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *BookController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.BookService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *BookController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	result := controller.BookService.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)

}
