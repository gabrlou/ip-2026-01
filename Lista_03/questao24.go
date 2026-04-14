package main

import (
	f "fmt"
	m "math"
)

func main() {
	var (
		A   float64
		sen float64
	)

	f.Printf("\nTabela de senos (aproximação pela série de Maclaurin):\n\n")
	f.Printf("Ângulo \"A\"\t\tsen(A)\n")

	for A = 0.0; A <= 6.3; A += 0.1 {

		sen = A
		sinal := -1.0
		exp := 3

		for i := 1; i <= 10; i++ { // Vão ser calculados 10 termos para aproximar o seno do ângulo A

			sen += sinal * (m.Pow(A, float64(exp))) / float64(fatorial(exp))

			sinal *= -1
			exp += 2
		}

		f.Printf("%.1f\t\t%.6f\n", A, sen)
	}

	f.Print("\n")
}

func fatorial(numero int) int {
	if numero == 0 {
		return 1
	}
	return numero * fatorial(numero-1)
}
