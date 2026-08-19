package helpers

import (
	"errors"
	"os"
)

func ExisteArchivo(ruta string) bool {
	_, err := os.Stat(ruta)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}
