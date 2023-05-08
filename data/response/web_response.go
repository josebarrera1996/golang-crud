package response

// Estableciendo el formato de respuesta. Será de tipo JSON
type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}
