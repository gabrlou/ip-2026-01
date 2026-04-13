package main

import (
	f "fmt"
)

func main() {
	var (
		N, i1, i2 int
	)

	f.Print("\nDigite o primeiro termo da série de Fetuccine: ")
	f.Scan(&i1)

	f.Print("\nDigite o segundo termo da série de Fetuccine: ")
	f.Scan(&i2)

	for {
		f.Print("\nDigite a quantidade (N >= 3) de termos da sequência de Fetuccine a serem exibidos: ")
		f.Scan(&N)

		if N < 3 {
			f.Print("\n	Número inválido! Digite um número N maior ou igual a 3.\n")
			continue
		} else {
			break
		}
	}

	penultimo := i1
	ultimo := i2
	proximo := 0

	f.Printf("\n%d %d", penultimo, ultimo)

	for i := 2; i < N; i++ {

		if i1 % 2 == 0 || i2 % 2 == 0 {
			proximo = ultimo - penultimo
		} else {
			proximo = ultimo + penultimo
		}

		penultimo = ultimo
		ultimo = proximo
		f.Printf(" %d", proximo)
	}
	f.Print("\n\n")
}
