package services

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
)

// Implementación de la lógica comercial
// Definiendo los métodos
type BookService interface {
	Create(ctx context.Context, request request.BookCreateRequest)
	Update(ctx context.Context, request request.BookUpdateRequest)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) response.BookResponse
	FindAll(ctx context.Context) []response.BookResponse
}
