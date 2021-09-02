package evento

// Un Evento representa una acción originada por un ente en una aplicación conectada a internet.
// Algunos ejemplos de eventos son:
//  Un inicio de sesión en un sitio web o aplicación móvil.
//  Una compra en una tienda en línea.
type Evento struct {
	// Grupo de entes al que pertenece el ente que origina el evento.
	GrupoDeEntes *GrupoDeEntes `json:"grupoDeEntes"`

	// Ente que origina el evento.
	Ente *Ente `json:"ente"`

	// Variables asociadas con el evento.
	Variables *Variables `json:"variables"`
}

// Un ResultadoDeEvento representa el resultado de la evaluación de fiabilidad de un evento.
type ResultadoDeEvento struct {
	// Identificador único del evento generado por Verifico.
	Identificador string `json:"identificador"`

	// Fiabilidad de que el evento haya sido originado por el ente que dice ser.
	// Es un número en el rango 0 <= fiabilidad < 1.
	// Normalmente una fiabilidad de 0,8 o superior indica que se puede confiar en el ente;
	// en caso contrario se debe solicitar una verificación al usuario.
	Fiabilidad float64 `json:"fiabilidad"`

	// Verificacion contiene información sobre la verificación del evento.
	Verificacion ResultadoDeEventoVerificacion `json:"verificacion"`
}

// Un ResultadoDeEventoVerificacion representa la información de verificación del evento.
type ResultadoDeEventoVerificacion struct {
	// Requerida indica si se requiere una verificación por parte del ente para poder validar el evento.
	Requerida bool `json:"requerida"`

	// FactoresDisponibles indica los factores disponibles para solicitar verificación al ente.
	// Los valores posibles son «sms» y «correoElectronico».
	FactoresDisponibles []string `json:"factoresDisponibles"`
}

// Nuevo retorna un evento construido a partir del identificador del grupo de entes, el identificador del ente, y las variables proporcionadas.
func Nuevo(grupoDeEntes string, ente string, variables *Variables) *Evento {
	return &Evento{
		GrupoDeEntes: &GrupoDeEntes{
			Identificador: grupoDeEntes,
		},
		Ente: &Ente{
			Identificador: ente,
		},
		Variables: variables,
	}
}
