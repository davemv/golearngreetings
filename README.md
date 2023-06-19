# Saludos en Go

Prueba para un paquete que proporciona una forma simple de obtener saludos en GO

## Instalación
Ejecutar:
```bash
go get -u github.com/davemv/golearngreetings

```

## Uso
package main

```go
import (
	"fmt"
	"log"

	greetings "github.com/davemv/golearngreetings"
)

func main() {
	log.SetPrefix("greetings: ")

	// el valor 0 hará que no se muestren ni la fecha ni hora
	log.SetFlags(5)
	// message, err := greetings.Hello("David")

	names := []string{"David", "Alex", "Roel", "Juan"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(message)
	fmt.Println(messages)
}
```