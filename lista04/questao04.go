package main

import f "fmt"

func main() {

	var numero int
	f.Print("\nDigite um número: ")
	f.Scan(&numero)

	if numero == 0 {
		f.Println("Número em binário: 0")
		return
	}

	f.Print("Número em binário: ")
	converter_binario(numero)
	f.Print("\n\n")
}

func converter_binario(numero int) {
	if numero == 0 {
		return
	}

	converter_binario(numero / 2)

	f.Print(numero % 2)
}
