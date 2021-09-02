package verifico

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/verifico/verifico-go/evento"
	"github.com/verifico/verifico-go/evento/verificacion"
)

// RegistrarEvento registra un nuevo evento recibido desde la aplicación cliente y obtiene el resultado de la evaluación de fiabilidad.
func (O *Cliente) RegistrarEvento(ctx context.Context, evto *evento.Evento) (*evento.ResultadoDeEvento, error) {
	if evto == nil {
		return nil, errors.New("evento nulo")
	}

	evto.GrupoDeEntes.Identificador = strings.TrimSpace(evto.GrupoDeEntes.Identificador)
	evto.Ente.Identificador = strings.TrimSpace(evto.Ente.Identificador)

	if evto.GrupoDeEntes.Identificador == "" {
		return nil, errors.New("falta el identificador del grupo de entes")
	}

	if evto.Ente.Identificador == "" {
		return nil, errors.New("falta el identificador del ente")
	}

	if evto.Variables == nil {
		return nil, errors.New("faltan las variables")
	}

	url := fmt.Sprintf("%s/eventos", URLAPI)
	cuerpo, err := O.peticion(ctx, "POST", url, evto)
	if err != nil {
		return nil, err
	}

	defer cuerpo.Close()

	var resultado = &evento.ResultadoDeEvento{}
	if err := json.NewDecoder(cuerpo).Decode(resultado); err != nil {
		return nil, err
	}

	return resultado, nil
}

// VerificarCodigoDeEvento verifica si el código proporcionado es válido para el evento dado.
func (O *Cliente) VerificarCodigoDeEvento(ctx context.Context, identificadorDeEvento string, codigo string) (*verificacion.ResultadoDeVerificacionDeEvento, error) {
	identificadorDeEvento = strings.TrimSpace(identificadorDeEvento)
	codigo = strings.TrimSpace(codigo)

	if identificadorDeEvento == "" {
		return nil, errors.New("falta el identificador del evento")
	}

	if codigo == "" {
		return nil, errors.New("falta el código a verificar")
	}

	vrfccn := &verificacion.VerificacionDeEvento{
		Codigo: codigo,
	}

	url := fmt.Sprintf("%s/eventos/%s/verificar-codigo", URLAPI, url.PathEscape(identificadorDeEvento))
	cuerpo, err := O.peticion(ctx, "POST", url, vrfccn)
	if err != nil {
		return nil, err
	}

	defer cuerpo.Close()

	var resultado = &verificacion.ResultadoDeVerificacionDeEvento{}
	if err := json.NewDecoder(cuerpo).Decode(resultado); err != nil {
		return nil, err
	}

	return resultado, nil
}
