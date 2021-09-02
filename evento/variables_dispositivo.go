package evento

// VariableDispositivo representa un dispositivo físico desde donde se origina un evento.
type VariableDispositivo struct {
	// Identificador único del dispositivo desde donde se origina el evento.
	// Normalmente es un número de serie, una dirección MAC, o un IMEI.
	Identificador string `json:"identificador,omitempty"`

	// Plataforma o sistema operativo del dispositivo desde donde se origina el evento.
	Plataforma string `json:"plataforma,omitempty"`

	// Metadatos del navegador web desde donde se origina el evento.
	Navegador *VariableNavegador `json:"navegador,omitempty"`
}

// VariableNavegador representa el navegador web desde donde se origina un envento.
type VariableNavegador struct {
	// Metadatos del objeto «window» del navegador.
	Window map[string]interface{} `json:"window,omitempty"`
}

// NuevaVariableDispositivo retorna una VariableDispositivo inicializada con las variables proporcioandas.
func NuevaVariableDispositivo(identificador string, plataforma string, navegador *VariableNavegador) *VariableDispositivo {
	return &VariableDispositivo{
		Identificador: identificador,
		Plataforma:    plataforma,
		Navegador:     navegador,
	}
}

// NuevaVariableNavegador retorna una VariableNavegador inicializada con los metadatos del objeto window del navegador proporcionados.
func NuevaVariableNavegador(window map[string]interface{}) *VariableNavegador {
	return &VariableNavegador{
		Window: window,
	}
}
