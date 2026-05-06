package main

import f "fmt"

func main() {

	vetor1 := make([]int, 10)
	vetor2 := make([]int, 10)
	resultante := make([]int, 20)

	f.Print("\n")
	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº número do vetor 1: ", i+1)
		f.Scan(&vetor1[i])
	}

	f.Print("\n")
	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº número do vetor 2: ", i+1)
		f.Scan(&vetor2[i])
	}

	j := 0
	for i := 0; i < 10; i++ {
		resultante[j] = vetor1[i]
		j++

		resultante[j] = vetor2[i]
		j++
	}

	f.Print("\nVetor resultante:\n")
	f.Print(resultante[:20])

	f.Print("\n\n")
}
