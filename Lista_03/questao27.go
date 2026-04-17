package main

import (
	f "fmt"
	m "math"
)

func main() {
	var (
		x   float64
		cos float64
	)

	f.Printf("\nDigite o valor de X para o cálculo de cos(x): ")
	f.Scan(&x)

	cos = 1
	sinal := -1.0
	exp := 2

	for i := 1; i <= 20; i++ { // Vão ser calculados 20 termos para aproximar o cos(x)

		cos += sinal * (m.Pow(x, float64(exp))) / fatorial(exp)

		sinal *= -1
		exp += 2
	}

	diferenca := cos - m.Cos(x)

	f.Printf("\n	Cos(x) calculado: %.20f\n	Diferença entre o valor calculado e o valor fornecido pela função: %.20f\n\n", cos, diferenca)
}

func fatorial(numero int) float64 {
	if numero == 0 {
		return 1
	}
	return float64(numero) * fatorial(numero-1)
}
