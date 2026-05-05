package main

import (
	f "fmt"
	m "math"
)

func main() {
    
	vetor := make([]float64, 15)
	var numero int

	f.Print("\n")
	for i := 0; i < 15; i++ {
		f.Printf("Digite o %dº número: ", i+1)
		f.Scan(&numero)

		if numero < 0 {
			vetor[i] = -1
		} else {
			vetor[i] = m.Sqrt(float64(numero))
		}
	}

	
	f.Print("\nValores armazenados:\n")
	f.Printf("%.2f", vetor)
	f.Print("\n\n")
}
