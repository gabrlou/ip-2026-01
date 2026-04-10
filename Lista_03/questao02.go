package main

import f "fmt"

func main() {
	var (
		soma, contador int
		media          float64
	)
	soma = 0
	contador = 0
	for i := 50; i <= 70; i += 2 {
		soma += i
		contador++
	}

	media = float64(soma) / float64(contador)

	f.Printf("\nFaixa de Análise: 50 a 70\nSoma de todos os números pares: %d\nMédia Aritmética: %.2f\n\n", soma, media)
}
