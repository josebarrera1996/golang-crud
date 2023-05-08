package repositories

import (
	"context"
	"golang-crud/models"
)

// Definiendo la interfaz con los respectivos métodos de un CRUD
type BookRepository interface {
	/* Definiendo los métodos */

	// Método para guardar un usuario
	Save(ctx context.Context, book models.Book)
	// Método para actualizar un usuario ya existente.
	Update(ctx context.Context, book models.Book)
	// Método para eliminar un usuario ya existente (gracias a su ID)
	Delete(ctx context.Context, bookId int)
	// Método para encontrar un usuario ya existente (gracias a su ID)
	FindById(ctx context.Context, bookId int) (models.Book, error)
	// Método para encontrar a todos los usuarios
	FindAll(ctx context.Context) []models.Book
}
