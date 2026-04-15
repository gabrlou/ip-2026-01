package main

import f "fmt"

func main() {

	f.Print("\nSoma de 10 números digitados pelo usuário.")
	
	var Array [10]float64
	for i := 0; i < len(Array); i++ {

		termo := 0.0
		f.Print("\nDigite o %d número a ser somado:", i+1)
		f.Scan(&termo)

		Array [i] = termo
	}

	f.Printf("Array: %v", Array)

	resultado := soma(Array)
	f.Printf("\nA soma da Array é: %.2f\n\n", resultado)
}

func soma (array[] float64) float64 {
		if len(array) == 0 {
			return 0
		}
		return array[0] + soma(array[1:])
	}
