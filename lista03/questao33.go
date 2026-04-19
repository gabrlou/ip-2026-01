package main

import f "fmt"

func main() {
	var n1, n2 int

	for {
		f.Print("\nDigite o valor de N1: ")
		f.Scan(&n1)

		if n1 <= 0 {
			f.Println("Número inválido. Digite um número maior que zero.")
			continue
		} else {
			break
		}
	}

	for {
		f.Print("Digite o valor de N2: ")
		f.Scan(&n2)

		if n2 <= 0 {
			f.Println("Número inválido. Digite um número maior que zero.")
			continue
		} else {
			break
		}
	}

	resto := n1
	quociente := 0

	for resto >= n2 {
		resto -= n2
		quociente++
	}

	f.Printf("\nResto (%d,%d) = %d\nQuociente (%d,%d) = %d\n\n", n1, n2, resto, n1, n2, quociente)
}
