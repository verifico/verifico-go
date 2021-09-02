package evento

// VariableOrigen representa el origen desde donde se origina un evento.
type VariableOrigen struct {
	DireccionIP string               `json:"direccionIP,omitempty"`
	Coordenadas *VariableCoordenadas `json:"coordenadas,omitempty"`
}

// VariableCoordenadas representa las coordenadas físicas desde donde se origina un evento.
type VariableCoordenadas struct {
	Latitud  float64 `json:"latitud,omitempty"`
	Longitud float64 `json:"longitud,omitempty"`
}

// NuevaVariableOrigen retorna una VariableOrigen inicializada con la dirección IP y coordenadas proporcionadas.
func NuevaVariableOrigen(direccionIP string, coordenadas *VariableCoordenadas) *VariableOrigen {
	return &VariableOrigen{
		DireccionIP: direccionIP,
		Coordenadas: coordenadas,
	}
}

// NuevaVariableCoordenadas retorna una VariableCoordenadas inicializada con las coordenadas proporcionadas.
func NuevaVariableCoordenadas(latitud float64, longitud float64) *VariableCoordenadas {
	return &VariableCoordenadas{
		Latitud:  latitud,
		Longitud: longitud,
	}
}
