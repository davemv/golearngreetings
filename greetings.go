package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un salud para la persona especificada
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("nombre vacío")
	}

	// devuelve un saludo que incluye el nombre en un mensaje
	// message := fmt.Sprintf("¡Hola, %v! ¡Bienvenido!", name)
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

// randomFormat devuelve uno de varios formatos de mensajes
// se selecciona de forma aleatoria
func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Qué bueno verte, %v!",
		"¡Saludo, %v! ¡Encantado de conocerte!",
	}

	return formats[rand.Intn(len(formats))]

}
