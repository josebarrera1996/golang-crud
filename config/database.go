package config

import (
	"database/sql"
	"fmt"
	"golang-crud/helper"

	_ "github.com/lib/pq" // Postgres goalng driver
	"github.com/rs/zerolog/log"
)

// Definiendo las propiedades de la conexión.
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbName   = "postgres"
)

// Definiendo la función para conectarse con la B.D de Postgres
func DatabaseConnection() *sql.DB {
	// Alojando en esta variable todas las propiedades definidas anteriormente con sus respectivos valores
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	// Abriendo la conexión
	db, err := sql.Open("postgres", sqlInfo)

	// Manejando los errores
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	// Logeando por consola
	log.Info().Msg("Connected to database.")

	return db
}
