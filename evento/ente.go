package evento

// Un Ente representa el ente que origina un evento.
// Un ente es un actor que interactúa en una aplicación conectada a internet.
// Normalmente es el usuario final que hace uso de una aplicación móvil, web, o de escritorio.
type Ente struct {
	// Identificador único del ente en base de datos.
	Identificador string `json:"identificador"`
}
