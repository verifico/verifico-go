package evento

import "strings"

// VariableAdquisicion representa una adquisición de un producto o servicio.
type VariableAdquisicion struct {
	Pago                  *VariablePago                  `json:"pago,omitempty"`
	ReceptorDeFacturacion *VariableReceptorDeFacturacion `json:"receptorDeFacturacion,omitempty"`
	DestinatarioDeEnvio   *VariableDestinatarioDeEnvio   `json:"destinatarioDeEnvio,omitempty"`
}

// VariablePago representa un pago realizado por la adquisición de un producto o servicio.
type VariablePago struct {
	Precio  *VariablePrecio        `json:"precio,omitempty"`
	Tarjeta *VariableTarjetaDePago `json:"tarjeta,omitempty"`
}

// VariableTarjetaDePago representa una tarjeta usada para realizar una compra.
type VariableTarjetaDePago struct {
	FechaDeExpiracion    string `json:"fechaDeExpiracion,omitempty"`
	UltimosCuatroDigitos string `json:"ultimosCuatroDigitos,omitempty"`
}

// VariablePrecio representa el monto pagado por una compra.
type VariablePrecio struct {
	Moneda string `json:"moneda,omitempty"`
	Monto  string `json:"monto,omitempty"`
}

// VariableReceptorDeFacturacion representa a la persona natural o jurídica a quien se emite la factura de una compra.
type VariableReceptorDeFacturacion struct {
	Nombre            string                   `json:"nombre,omitempty"`
	Identificacion    string                   `json:"identificacion,omitempty"`
	NumeroTelefonico  string                   `json:"numeroTelefonico,omitempty"`
	CorreoElectronico string                   `json:"correoElectronico,omitempty"`
	DireccionPostal   *VariableDireccionPostal `json:"direccionPostal,omitempty"`
}

// VariableDestinatarioDeEnvio representa a la persona natural o jurídica a quien se envía un producto adquirido.
type VariableDestinatarioDeEnvio struct {
	Nombre            string                   `json:"nombre,omitempty"`
	NumeroTelefonico  string                   `json:"numeroTelefonico,omitempty"`
	CorreoElectronico string                   `json:"correoElectronico,omitempty"`
	DireccionPostal   *VariableDireccionPostal `json:"direccionPostal,omitempty"`
}

// VariableDireccionPostal representa una dirección postal.
type VariableDireccionPostal struct {
	Pais                      string `json:"pais,omitempty"`
	SubdivisionDePrimerNivel  string `json:"subdivisionDePrimerNivel,omitempty"`
	SubdivisionDeSegundoNivel string `json:"subdivisionDeSegundoNivel,omitempty"`
	SubdivisionDeTercerNivel  string `json:"subdivisionDeTercerNivel,omitempty"`
	DireccionFisica           string `json:"direccionFisica,omitempty"`
}

// NuevaVariableAdquisicion retorna una VariableAdquisicion inicializada con las variables proporcionadas.
func NuevaVariableAdquisicion(pago *VariablePago, receptorDeFacturacion *VariableReceptorDeFacturacion, destinatarioDeEnvio *VariableDestinatarioDeEnvio) *VariableAdquisicion {
	return &VariableAdquisicion{
		Pago:                  pago,
		ReceptorDeFacturacion: receptorDeFacturacion,
		DestinatarioDeEnvio:   destinatarioDeEnvio,
	}
}

// NuevaVariablePago retorna una VariablePago inicializada con los datos de precio y tarjeta proporcionados.
func NuevaVariablePago(precio *VariablePrecio, tarjeta *VariableTarjetaDePago) *VariablePago {
	return &VariablePago{
		Precio:  precio,
		Tarjeta: tarjeta,
	}
}

// NuevaVariableTarjetaDePago retorna una VariableTarjetaDePago inicializada con la fecha de expiración y los últimos cuatro dígitos proporcionados.
func NuevaVariableTarjetaDePago(fechaDeExpiracion string, ultimosCuatroDigitos string) *VariableTarjetaDePago {
	ultimosCuatroDigitos = strings.TrimSpace(ultimosCuatroDigitos)
	var l = len(ultimosCuatroDigitos)
	if l > 4 {
		ultimosCuatroDigitos = ultimosCuatroDigitos[l-4:]
	}

	return &VariableTarjetaDePago{
		FechaDeExpiracion:    fechaDeExpiracion,
		UltimosCuatroDigitos: ultimosCuatroDigitos,
	}
}

// NuevaVariablePrecio retorna una VariablePrecio inicializada con la moneda y el monto proporcionados.
func NuevaVariablePrecio(moneda string, monto string) *VariablePrecio {
	return &VariablePrecio{
		Moneda: moneda,
		Monto:  monto,
	}
}

// NuevaVariableReceptorDeFacturacion retorna una VariableReceptorDeFacturacion inicializada con las variables proporcionadas.
func NuevaVariableReceptorDeFacturacion(nombre, identificacion, numeroTelefonico, correoElectronico string, direccionPostal *VariableDireccionPostal) *VariableReceptorDeFacturacion {
	return &VariableReceptorDeFacturacion{
		Nombre:            nombre,
		Identificacion:    identificacion,
		NumeroTelefonico:  numeroTelefonico,
		CorreoElectronico: correoElectronico,
		DireccionPostal:   direccionPostal,
	}
}

// NuevaVariableDestinatarioDeEnvio retorna una VariableDestinatarioDeEnvio inicializada con las variables proporcionadas.
func NuevaVariableDestinatarioDeEnvio(nombre, numeroTelefonico, correoElectronico string, direccionPostal *VariableDireccionPostal) *VariableDestinatarioDeEnvio {
	return &VariableDestinatarioDeEnvio{
		Nombre:            nombre,
		NumeroTelefonico:  numeroTelefonico,
		CorreoElectronico: correoElectronico,
		DireccionPostal:   direccionPostal,
	}
}

// NuevaVariableDireccionPostal retorna una VariableDireccionPostal inicializada con las variables proporcionadas.
func NuevaVariableDireccionPostal(pais, subdivisionDe1erNivel, subdivisionDe2doNivel, subdivisionDe3erNivel, direccionFisica string) *VariableDireccionPostal {
	return &VariableDireccionPostal{
		Pais:                      pais,
		SubdivisionDePrimerNivel:  subdivisionDe1erNivel,
		SubdivisionDeSegundoNivel: subdivisionDe2doNivel,
		SubdivisionDeTercerNivel:  subdivisionDe3erNivel,
		DireccionFisica:           direccionFisica,
	}
}
