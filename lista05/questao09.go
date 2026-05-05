package main

import f "fmt"

func main() {
    
	alturas := make([]float64, 10)
	var soma float64

	f.Print("\n")
	for i := 0; i < 10; i++ {
		f.Printf("Digite a altura do %dº atleta: ", i+1)
		f.Scan(&alturas[i])
		soma += alturas[i]
	}
	
	media := soma / 10

	f.Printf("\nMédia das alturas: %.2f\n", media)
	f.Print("Alturas acima da média:\n")

	for i := 0; i < 10; i++ {
		if alturas[i] > media {
			f.Printf("%.2f ", alturas[i])
		}
	}

	f.Print("\n\n")
}
