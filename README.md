# github.com/verifico/verifico-go

**Cliente Go para la API de prevención de suplantación de identidad de Verifico.**

## Uso


**Ejemplo de uso de registro de evento:**

```go
package main

import (
	"context"

	verifico "github.com/verifico/verifico-go"
	"github.com/verifico/verifico-go/evento"
)

func main() {

	llave := "<LLAVE_DE_AUTORIZACIÓN_DE_API>"
	grupoDeEntes := "<IDENTIFICADOR_DE_GRUPO_DE_ENTES>"
	ente := "<IDENTIFICADOR_DE_ENTE>"

	varDireccionIP := "1.2.3.4"
	varCoordenadas := evento.NuevaVariableCoordenadas(4.1234, -72.4321)
	varOrigen := evento.NuevaVariableOrigen(varDireccionIP, varCoordenadas)

	varNavegador := evento.NuevaVariableNavegador(map[string]interface{}{
		"innerHeight": 651,
		"innerWidth":  1880,
	})
	varDispositivo := evento.NuevaVariableDispositivo("hfdy8w45", "Linux x86_64", varNavegador)

	varPrecio := evento.NuevaVariablePrecio("COP", "64000")
	varTarjetaDePago := evento.NuevaVariableTarjetaDePago("12/24", "9876")
	varPago := evento.NuevaVariablePago(varPrecio, varTarjetaDePago)
	varDireccionPostal := evento.NuevaVariableDireccionPostal("CO", "Bogotá", "Suba", "", "Calle 1 # 2-3")
	varReceptor := evento.NuevaVariableReceptorDeFacturacion("Juan Pérez", "1234567890", "3001234567", "juan@ejemplo.co", varDireccionPostal)
	varDestinatario := evento.NuevaVariableDestinatarioDeEnvio("Juan Pérez", "3001234567", "juan@ejemplo.co", varDireccionPostal)
	varAdquisicion := evento.NuevaVariableAdquisicion(varPago, varReceptor, varDestinatario)

	varOtras := map[string]interface{}{
		"artículoDeLujo":     true,
		"enPromoción":        true,
		"intencionesPrevias": 2,
		// Otras variables personalizadas...
	}

	variables := evento.NuevasVariables(varOrigen, varDispositivo, varAdquisicion, varOtras)

	evto := evento.Nuevo(grupoDeEntes, ente, variables)
	cliente := verifico.Nuevo(llave)

	resultado, err := cliente.RegistrarEvento(context.TODO(), evto)
	if err != nil {
		// Manejar error
		return
	}

	verificado := !resultado.Verificacion.Requerida
	if verificado {
		// Proceder con la compra.
	} else {
		// Solicitar verificación de código al usuario...
	}
}

```


**Ejemplo de uso de verificación de evento mediante código enviado al usuario:**

```go
package main

import (
	"context"

	verifico "github.com/verifico/verifico-go"
)

func main() {
	llave := "<LLAVE_DE_AUTORIZACIÓN_DE_API>"
	identificadorDeEvento := "<IDENTIFICADOR_DE_EVENTO_NO_VERIFICADO>"
	codigo := "<CÓDIGO_SOLICITADO_AL_USUARIO>"

	cliente := verifico.Nuevo(llave)
	verificacion, err := cliente.VerificarCodigoDeEvento(context.TODO(), identificadorDeEvento, codigo)
	if err != nil {
		// Manejar error
		return
	}

	if verificacion.Valido {
		// Proceder con la compra.
	} else {
		// Impedir compra.
	}
}
```