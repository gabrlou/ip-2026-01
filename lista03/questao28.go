package main

import (
	f "fmt"
	m "math"
)

func main() {
	var S, pi, sinal float64

	S = 0
	sinal = 1.0
	denominador := 1

	for i := 1; i <= 51; i++ { // Vão ser calculados 51 termos para aproximar o valor de π (pi)

		S += sinal * 1 / m.Pow(float64(denominador), 3)

		sinal *= -1
		denominador += 2
	}

	pi = m.Pow(S*32, 1.0/3.0)

	f.Printf("\n	Valor de π (pi): %.15f\n\n", pi)
}
