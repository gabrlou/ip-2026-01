package main

import (
	"fmt"
	f "fmt"
)

func main() {

	f.Print("\n	Índices dos elementos de uma matriz 10x10 acima da diagonal principal:\n\n")
	for i := 0; i < 10; i++ { // Linhas: "i"
		for j := 0; j < 10; j++ { // Colunas: "j"
			if j > i {
				fmt.Printf("(%d, %d) ", i, j)
				continue
			}
		}
		fmt.Println()
	}
	f.Print("\n")
}
