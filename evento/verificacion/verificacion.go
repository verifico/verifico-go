package verificacion

// Una VerificacionDeEvento representa una verificación requerida por un evento previo.
type VerificacionDeEvento struct {
	Codigo string `json:"codigo"`
}

// Un ResultadoDeVerificacionDeEvento representa el resultado de la verificación requerida por un evento previo.
type ResultadoDeVerificacionDeEvento struct {
	Valido bool `json:"valido"`
}
