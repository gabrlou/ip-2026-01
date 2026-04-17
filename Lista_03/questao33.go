package main

import f "fmt"

func main() {
	var n1, n2 int

	f.Print("\nDigite o valor de N1: ")
	f.Scan(&n1)

	f.Print("Digite o valor de N2: ")
	f.Scan(&n2)

	resto := n1
	quociente := 0

	for resto >= n2 {
		resto -= n2
		quociente++
	}

	f.Printf("\nResto (%d,%d) = %d\nQuociente (%d,%d) = %d\n\n", n1, n2, resto, n1, n2, quociente)
}
