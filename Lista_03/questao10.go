package main

import (
	f "fmt"
)

func main() {
	var (
		N int
	)

	for {
		f.Print("\nDigite a quantidade (N >= 3) de termos da sequência de Fibonacci a serem exibidos: ")
		f.Scan(&N)

		if N < 3 {
			f.Print("\n	Número inválido! Digite um número N maior ou igual a 3.\n")
			continue
		} else {
			break
		}
	}

	penultimo := 0
	ultimo := 1
	f.Printf("%d - %d", penultimo, ultimo)

	for i := 2; i < N; i++ {
		proximo := penultimo + ultimo
		penultimo = ultimo
		ultimo = proximo
		f.Printf(" - %d", proximo)
	}
	f.Print("\n\n")
}
