package main

import f "fmt"

func main() {

	numerador := 1
	denominador := 15
	soma := 0.0

	for i := 1; denominador >= 1; i++ {

		sinal := 1.0
		if i%2 == 0 {
			sinal = -1.0
		}

		termo := sinal * (float64(numerador) / (float64(denominador * denominador)))
		soma += termo
		numerador *= 2
		denominador -= 1
	}
	f.Printf("\nResultado da soma [S = 1/225 - 2/196 + 4/169 ...]: %.2f\n\n", soma)
}
