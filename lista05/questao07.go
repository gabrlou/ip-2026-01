package main

import f "fmt"

func main() {
	vetor := make([]int, 100)

	for i := 0; i < 100; i++ {
		vetor[i] = 2*i + 1
	}

	f.Print(vetor[:100])
	f.Print("\n\n")
}
