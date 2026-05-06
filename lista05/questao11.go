package main

import (
	f "fmt"
	m "math"
)

func main() {

	B := make([]float64, 100)
	var S float64

	f.Print("\n")
	for i := 0; i < 100; i++ {
		f.Printf("Digite o %dº valor: ", i+1)
		f.Scan(&B[i])
	}

	for i := 0; i < 50; i++ {
		diferenca := B[i] - B[99-i]
		S += m.Pow(diferenca, 3)
	}

	f.Printf("\nValor de S: %.2f\n\n", S)
}
