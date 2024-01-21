# Saludos en Go
Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Instalacion
Ejecuta el siguiente para instalar el paquete:
```bash
go get -u github.com/ValenOl/greetings
```

## Uso
Ejemplo de como utilizar el paquete en tu codigo:

```go
package main

import (
	"fmt"
	"log"

	"github.com/ValenOl/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Valentin", "Messi", "De Paul"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	/*message, err := greetings.Hello("Valentin")
	if err != nil {
		log.Fatal(err)
	} */

	fmt.Println(messages)
}

```

Este ejemplo importa el paquete github.com/ValenOl/greetings y llama a la funcion de hello.
