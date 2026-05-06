package main

import f "fmt"

func main() {

	var num int
	vetor := make([]int, 10)

	f.Print("\n")
	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº número: ", i+1)
		f.Scan(&num)

		j := i - 1
		for j >= 0 && vetor[j] > num {
			vetor[j+1] = vetor[j]
			j--
		}

		vetor[j+1] = num
	}

	f.Print("\nVetor ordenado:\n")
	f.Print(vetor[:10])
	f.Print("\n\n")
}
