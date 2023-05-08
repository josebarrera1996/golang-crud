package helper

// Helper para poder trabajar con los errores
func PanicIfError(err error) {
	if err != nil {
		// En caso de error, detener la ejecución del programa
		panic(err)
	}
}
