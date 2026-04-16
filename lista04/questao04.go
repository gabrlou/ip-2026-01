package main

import (
	"fmt"
	f "fmt"
)

func main() {

	var numero int
	fmt.Print("\nDigite um número: ")
	fmt.Scan(&numero)

	if numero == 0 {
		fmt.Println("Número em binário: 0")
		return
	}

	fmt.Print("Número em binário: ")
	converter_binario(numero)
	f.Print("\n\n")
}

func converter_binario(numero int) {
	if numero == 0 {
		return
	}

	converter_binario(numero / 2)

	fmt.Print(numero % 2)
}
