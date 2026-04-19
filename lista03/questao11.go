package main

import f "fmt"

func main() {
	var N, resultado int

	for {
		f.Print("\nDigite um número inteiro para calcular sua fatorial: ")
		f.Scan(&N)
		if N < 0 {
			f.Print("\nNúmero inválido! Digite um número maior ou igual a zero.\n")
			continue
		} else if N == 0 {
			resultado = 1
			break
		} else {
			resultado = 1
			for i := 1; i <= N; i++ {
				resultado *= i
			}
			break
		}
	}
	f.Printf("Resultado: %d! = %d\n\n", N, resultado)
}
