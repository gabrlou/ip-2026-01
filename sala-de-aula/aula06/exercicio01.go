package main

import f "fmt"

func main() {
	const quant_notas int = 3 // Quantidade de notas
	var (
		nota  [quant_notas]float64 // Array de notas
		soma  float64              = 0
		media float64
	)

	for i := 0; i < quant_notas; i++ {
		f.Printf("Informe a %dº nota : ", i+1)
		f.Scan(&nota[i])
		soma += nota[i]
	}
	media = soma / float64(quant_notas)
	f.Printf("\nMédia geral: %.2f", media)

	f.Print("\nNotas acima da média: ")
	for i := 0; i < quant_notas; i++ {
		if nota[i] > media {
			f.Print(nota[i], "; ")
		}

	}
}
