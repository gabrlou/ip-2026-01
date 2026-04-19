package main

import f "fmt"

func main() {

	f.Print("\nSoma de 10 números digitados pelo usuário.\n")

	var Array [10]float64

	for i := 0; i < len(Array); i++ {
		f.Printf("Digite o %dº número a ser somado: ", i+1)

		var termo float64
		f.Scan(&termo)

		Array[i] = termo
	}

	f.Printf("\nArray: %v\n", Array)

	resultado := soma(Array[:])
	f.Printf("A soma do Array é: %.2f\n", resultado)
}

func soma(array []float64) float64 {
	if len(array) == 0 {
		return 0
	}
	return array[0] + soma(array[1:])
}
