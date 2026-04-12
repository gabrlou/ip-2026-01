package main

import f "fmt"

func main() {
	var soma = 0

	f.Print("\n")
	for i := 1; i <= 20; i++ {
		f.Printf("%d ", i)
		soma += i
	}
	f.Printf("\nSoma: %d\n\n", soma)
}
