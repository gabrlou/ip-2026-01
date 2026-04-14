package main

import (
	"fmt"
	f "fmt"
)

func main() {

	f.Print("\n	Índices de todos os elementos de uma matriz 10x10:\n\n")
	for i := 0; i < 10; i++ { // Linhas: "i"
		for j := 0; j < 10; j++ { // Colunas: "j"
			fmt.Printf("(%d, %d) ", i, j)
		}
		fmt.Println()
	}
	f.Print("\n")
}
