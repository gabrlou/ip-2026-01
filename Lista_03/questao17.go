package main

import f "fmt"

func main() {

	f.Print("\n	Índices de todos os elementos de uma matriz 10x10:\n\n")
	for i := 0; i < 10; i++ { // Linhas: "i"
		for j := 0; j < 10; j++ { // Colunas: "j"
			f.Printf("(%d, %d) ", i, j)
		}
		f.Println()
	}
	f.Print("\n")
}
