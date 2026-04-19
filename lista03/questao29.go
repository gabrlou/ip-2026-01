package main

import f "fmt"

func main() {
	var N, soma int

	for {
		f.Print("\nDigite o valor de N para obter a somatória de 1 a N: ")
		f.Scan(&N)
		if N <= 0 {
			f.Print("\nNúmero inválido. Digite um valor maior que zero.\n\n")
		} else {
			break
		}
	}

	soma = 0

	for i := 1; i <= N; i++ {
		soma += i
	}

	f.Printf("\n	Soma de 1 a %d: %d\n\n", N, soma)
}
