package verifico

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// Un Cliente representa una instancia para realizar operaciones con la API de Verifico.
type Cliente struct {
	llave string
}

// Nuevo retorna un cliente inicializado con la llave de conexión a la API de Verifico.
func Nuevo(llave string) *Cliente {
	return &Cliente{
		llave: llave,
	}
}

// peticion genera una petición HTTP a la API de Verifico.
func (O *Cliente) peticion(ctx context.Context, metodo string, recurso string, datos interface{}) (io.ReadCloser, error) {
	if O.llave == "" {
		return nil, errors.New("falta la llave de autenticación")
	}

	datosJSON, err := json.Marshal(datos)
	if err != nil {
		return nil, err
	}

	C := &http.Client{
		Timeout: 20 * time.Second,
	}

	R, err := http.NewRequestWithContext(ctx, metodo, recurso, bytes.NewBuffer(datosJSON))
	if err != nil {
		return nil, err
	}

	R.Header.Set("Content-Type", "application/json")
	R.Header.Set("Authorization", O.llave)

	res, err := C.Do(R)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		defer res.Body.Close()

		er := "[8888] ocurrió un error"
		resultado := map[string]string{}
		json.NewDecoder(res.Body).Decode(&resultado)
		if resultado["error"] != "" {
			er = resultado["error"]
		}

		return nil, errors.New(er)
	}

	return res.Body, nil
}
