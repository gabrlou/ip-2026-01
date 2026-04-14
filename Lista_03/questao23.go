package main

import f "fmt"

func main() {
	var (
		N int
	)

	for {
		f.Print("\nDigite a quantidade N de termos da série [1000/1 - 997/2 + 994/3 ...] a serem calculados: ")
		f.Scan(&N)
		if N <= 0 {
			f.Print("\nNúmero inválido. Digite um valor maior que zero.\n")
			continue
		} else {
			break
		}
	}

	numerador := 1000.0
	soma := 0.0

	for i := 1; i <= N; i++ {

		sinal := 1.0
		if i%2 == 0 {
			sinal = -1.0
		}

		termo := sinal * (numerador / float64(i))
		soma += termo
		numerador -= 3
	}
	f.Printf("\nTotal de termos calculados: %d\nResultado da soma: %.2f\n\n", N, soma)
}
