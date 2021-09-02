package evento

// Variables representa las variables recolectadas de un evento.
type Variables struct {
	Origen      *VariableOrigen      `json:"origen,omitempty"`
	Dispositivo *VariableDispositivo `json:"dispositivo,omitempty"`
	Adquisicion *VariableAdquisicion `json:"adquisicion,omitempty"`

	// Variables personalizadas asociadas a un evento.
	// El valor del mapa debe ser un tipo básico: bool, string, (u)int*, float*.
	// Cualquier otro valor se omitirá.
	Otras map[string]interface{} `json:"otras,omitempty"`
}

// NuevasVariables retorna el objeto Variables inicializado con las variables proporcionadas.
func NuevasVariables(origen *VariableOrigen, dispositivo *VariableDispositivo, adquisicion *VariableAdquisicion, otras map[string]interface{}) *Variables {
	return &Variables{
		Origen:      origen,
		Dispositivo: dispositivo,
		Adquisicion: adquisicion,
		Otras:       otras,
	}
}
