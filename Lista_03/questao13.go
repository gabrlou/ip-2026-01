package main

import f "fmt"

func main() {
	var (
		H, termo, numerador, denominador float64
	)

	H = 0.0
	numerador = 1
	denominador = 1
	for i := 0; numerador <= 99; i++ {
		termo = numerador / denominador
		numerador += 2
		denominador++
		H += termo
	}
	f.Printf("\n	Resultado da soma [H = 1/1 + 3/2 + 5/3 ... 99/50]: %f\n\n", H)
}
