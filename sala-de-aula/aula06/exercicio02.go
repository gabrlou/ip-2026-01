package main

import f "fmt"

func main() {
	var (
		numeros [5]int
		soma    int = 0
	)

	for i := 0; i < len(numeros); i++ {
		f.Printf("Informe o %dº número: ", i+1)
		f.Scan(&numeros[i])
		soma += numeros[i]
	}

	f.Printf("\nA soma dos 5 valores é: %d\n", soma)
}
