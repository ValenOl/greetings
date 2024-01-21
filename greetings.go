package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// HELLO - Devuelve un saludo para la persona especificada

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Nombre vacio")
	}

	// Devuelve un saludo que incluye el nombre en un mensaje
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// randomFormat devuelve uno de varios formatos de de saludo
// Se selecciona de forma aleatoria
func randomFormat() string {
	formats := []string{
		"Hola, %v! Bienvenido!",
		"Un gusto verte de nuevo por aca, %v",
		"%v!!! Que sorpresa por aca, nuevamente.",
	}

	return formats[rand.Intn(len(formats))]
}
