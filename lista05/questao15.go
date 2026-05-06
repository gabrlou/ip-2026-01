package main

import f "fmt"

func main() {

	vetor1 := make([]int, 30)
	vetor2 := make([]int, 30)

	f.Print("\n")
	for i := 0; i < 30; i++ {
		f.Printf("Digite o %dº número: ", i+1)
		f.Scan(&vetor1[i])
	}

	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			vetor2[i] = vetor1[i] * 2
		} else {
			vetor2[i] = vetor1[i] * 3
		}
	}

	f.Print("\nVetor resultante:\n")
	f.Print(vetor2[:30])
	f.Print("\n\n")
}
