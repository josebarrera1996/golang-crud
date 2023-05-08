package repositories

import (
	"context"
	"database/sql"
	"errors"
	"golang-crud/helper"
	"golang-crud/models"
)

// Utilizando los métodos de la librería SQL
type BookRepositoryImpl struct {
	Db *sql.DB
}

// Preparando para la implementación de los métodos definidos en BookRepository
func NewBookRepository(Db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

/* Implementando los métodos del repositorio */

// Implementando el método de Delete
func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// QUery
	SQL := "delete from book where id =$1"
	_, errExec := tx.ExecContext(ctx, SQL, bookId)
	helper.PanicIfError(errExec)
}

// Implementando el método para encontrar a los libros
func (b *BookRepositoryImpl) FindAll(ctx context.Context) []models.Book {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// query
	SQL := "select id,name from book"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var books []models.Book

	for result.Next() {
		book := models.Book{}
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)

		books = append(books, book)
	}

	return books
}

// Implementando el método para encontrar a un libro en específico
func (b *BookRepositoryImpl) FindById(ctx context.Context, bookId int) (models.Book, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id,name from book where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	book := models.Book{}

	if result.Next() {
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)
		return book, nil
	} else {
		return book, errors.New("book id not found")
	}
}

// Implementando el método para guardar un libro
func (b *BookRepositoryImpl) Save(ctx context.Context, book models.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into book(name) values ($1)"
	_, err = tx.ExecContext(ctx, SQL, book.Name)
	helper.PanicIfError(err)
}

// Implementando el método para actualizar a un libro en específico
func (b *BookRepositoryImpl) Update(ctx context.Context, book models.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update book set name=$1 where id=$2"
	_, err = tx.ExecContext(ctx, SQL, book.Name, book.Id)
	helper.PanicIfError(err)
}
