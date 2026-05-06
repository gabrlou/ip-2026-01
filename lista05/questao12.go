package main

import f "fmt"

func main() {

	notas := make([]int, 15)
	freqAbs := make([]int, 11)

	f.Print("\n")
	for i := 0; i < 15; i++ {
		f.Printf("Digite a %dª nota (0 a 10): ", i+1)
		f.Scan(&notas[i])

		freqAbs[notas[i]]++
	}

	f.Print("\nNota\tFrequência Absoluta\tFrequência Relativa\n")

	for i := 0; i <= 10; i++ {
		freqRel := float64(freqAbs[i]) / 15.0
		f.Printf("%d\t\t%d\t\t\t%.2f\n", i, freqAbs[i], freqRel)
	}

	f.Print("\n\n")
}
