package main

import f "fmt"

func main() {
	jogadas := make([]int, 20)
	freq := make([]int, 7)

	f.Print("\n")
	for i := 0; i < 20; i++ {
		f.Printf("Digite o resultado da %dª jogada (1 a 6): ", i+1)
		f.Scan(&jogadas[i])

		if jogadas[i] >= 1 && jogadas[i] <= 6 {
			freq[jogadas[i]]++
		}
	}

	f.Print("\nJogadas:\n")
	f.Print(jogadas[:20])

	f.Print("\nFrequência:\n")
	for i := 1; i <= 6; i++ {
		f.Printf("Face %d apareceu %d vez(es)\n", i, freq[i])
	}

	f.Print("\n")
}
