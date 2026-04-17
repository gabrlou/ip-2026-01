package main

import "fmt"

func main() {
	var graos, total uint64

	graos = 1
	total = 0

	for i := 1; i <= 64; i++ {
		total += graos
		graos *= 2
	}

	fmt.Printf("\nTotal de grãos no tabuleiro: %d\n\n", total)
}
