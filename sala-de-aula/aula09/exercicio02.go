package main

import f "fmt"

func main() {
	var (
		indice int
		x int
	)

	lista := []int{1, 5, 11, 15, 20, 21, 24, 45, 67, 76}

	f.Print("\nDigite o valor a ser procurado na lista: ")
	f.Scan(&x)

	indice = BuscaBinaria(lista, x)

	if indice != -1 {
		f.Printf("\nValor encontrado no índice: %d\n\n", indice)
	} else {
		f.Print("\nValor não encontrado na lista.\n\n")
	}
}

func BuscaBinaria(l[]int, x int) int {

	n := len(l)
	pos_ini := 0
	pos_fim := n - 1

	for pos_ini <= pos_fim {

		pos_meio := (pos_ini + pos_fim) / 2

		if l[pos_meio] == x {
			return pos_meio
		} else if l[pos_meio] < x {
			pos_ini = pos_meio + 1
		} else if l[pos_meio] > x {
			pos_fim = pos_meio - 1
		}
	}
	return -1
}
