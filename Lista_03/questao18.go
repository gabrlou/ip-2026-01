package main

import f "fmt"

func main() {

	f.Print("\n	Índices da diagonal principal de uma matriz 10x10:\n\n")
	for i := 0; i < 10; i++ { // Linhas: "i"
		for j := 0; j < 10; j++ { // Colunas: "j"
			if i == j {
				fmt.Printf("(%d, %d) ", i, j)
				continue
			}
		}
		fmt.Println()
	}
	f.Print("\n")
}
