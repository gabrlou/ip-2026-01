package main

import f "fmt"

func main() {
	var numero, dezena int

	f.Print("\nDigite um número inteiro de 3 casas: ")
	f.Scan(&numero)

	if numero >= 100 && numero <= 999 {
		dezena = (numero / 10) % 10
		f.Printf("Dezena(s): %d\n", dezena)
	} else {
		f.Println("Número inválido. Digite um número de 3 casas.\n")
	}
}
