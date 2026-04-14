package main

import f "fmt"

func main() {
	var (
		indice int
		x int
	)

	lista := []int{20, 5, 15, 24, 67, 45, 1, 76, 21, 11}

	f.Print("\nDigite o valor a ser procurado na lista: ")
	f.Scan(&x)

	indice = BuscaSequencial(lista, x)

	if indice != -1 {
		f.Printf("\nValor encontrado no índice: %d\n\n", indice)
	} else {
		f.Print("\nValor não encontrado na lista.\n\n")
	}
	
}

func BuscaSequencial(l[]int, x int) int {
	for i:= 0; i < len(l); i++ {
		if l[i] == x {
			return i
		}
	}
	return -1
}
