package main

import f "fmt"

func main() {

	soma := 0.0

	for i := 0; i < 100; i++ {

		numerador := 100 - i
		soma += float64(numerador) / Fatorial(i)
	}

	f.Printf("\nResultado da soma [S = 100/0! + 99/1! + 98/2! ...]: %.2f\n\n", soma)
}

func Fatorial(numero int) float64 {
	resultado := 1.0

	for i := 2; i <= numero; i++ {
		resultado *= float64(i)
	}

	return resultado
}
