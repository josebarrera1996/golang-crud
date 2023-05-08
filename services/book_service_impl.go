package services

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/models"
	"golang-crud/repositories"
)

// Implementando los métodos definidos en el repositorio, utilizando el BookService
type BookServiceImpl struct {
	BookRepository repositories.BookRepository
}

func NewBookServiceImpl(bookRepository repositories.BookRepository) BookService {
	return &BookServiceImpl{BookRepository: bookRepository}
}

/* Implementando los métodos */

// Implementación para crear un nuevo libro
func (b *BookServiceImpl) Create(ctx context.Context, request request.BookCreateRequest) {
	book := models.Book{
		Name: request.Name,
	}
	b.BookRepository.Save(ctx, book)
}

// Implementación para borrar un libro
func (b *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helper.PanicIfError(err)
	b.BookRepository.Delete(ctx, book.Id)
}

// Implementación para encontrar a todos los libros
func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
	books := b.BookRepository.FindAll(ctx)

	var bookResp []response.BookResponse

	for _, value := range books {
		book := response.BookResponse{Id: value.Id, Name: value.Name}
		bookResp = append(bookResp, book)
	}
	return bookResp

}

// Implementación para encontrar a un libro en especifico
func (b *BookServiceImpl) FindById(ctx context.Context, bookId int) response.BookResponse {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helper.PanicIfError(err)
	return response.BookResponse(book)
}

// Implementación para actualizar a un libro en específico
func (b *BookServiceImpl) Update(ctx context.Context, request request.BookUpdateRequest) {
	book, err := b.BookRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	book.Name = request.Name
	b.BookRepository.Update(ctx, book)
}
