package main

import f "fmt"

func main() {

	vetor := make([]int, 10)

	f.Print("\n")
	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº número: ", i+1)
		f.Scan(&vetor[i])
	}

	X := vetor[0]
	P := 0

	for i := 1; i < 10; i++ {
		if vetor[i] < X {
			X = vetor[i]
			P = i
		}
	}

	f.Printf("\nO menor elemento do vetor é %d e sua posição dentro do vetor é: %d\n\n", X, P)
}
