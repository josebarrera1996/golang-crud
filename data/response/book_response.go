package response

// Estableciendo el formato de la respuesta. Será de tipo JSON
type BookResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
