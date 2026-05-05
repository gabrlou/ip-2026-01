package main

import f "fmt"

func main() {

	vetor := make([]int, 100)
	j := 0

	for i := 100; i >= 1; i-- {
		vetor[j] = i
		j++
	}

	f.Print("Vetor:\n")
	for i := 0; i < 100; i++ {
		f.Printf("%d ", vetor[i])
	}

	f.Print("\n\n")
}