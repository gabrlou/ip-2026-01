package main

import f "fmt"

func main() {
	var (
		X, soma float64
	)

	for {
		f.Print("\nDigite um número real X para calcular a soma [S = X - X/1! + X/2! ...]: ")
		f.Scan(&X)

		if X <= 0 {
			f.Print("\nNúmero inválido! Digite um número maior que zero.\n")
			continue
		}
		break
	}

	soma = 0.0
	for i := 0; i < 20; i++ {

		resultado := 1.0
		for j := 1; j <= i; j++ {
			resultado *= float64(j)
		}

		sinal := 1.0
		if i%2 == 1 {
			sinal = -1.0
		}

		termo := sinal * (X / resultado)
		soma += termo
	}

	f.Printf("	Resultado da soma: %f\n\n", soma)
}
