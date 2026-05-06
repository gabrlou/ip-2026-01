package main

import f "fmt"

func main() {

	v1 := make([]int, 10)
	v2 := make([]int, 5)

	f.Print("\n")
	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº número do vetor 1: ", i+1)
		f.Scan(&v1[i])
	}

	f.Print("\n")
	for i := 0; i < 5; i++ {
		f.Printf("Digite o %dº número do vetor 2: ", i+1)
		f.Scan(&v2[i])
	}

	f.Print("\n")
	for i := 0; i < 10; i++ {
		f.Printf("Número %d:\n", v1[i])

		encontrou := false

		for j := 0; j < 5; j++ {
			if v2[j] != 0 && v1[i]%v2[j] == 0 {
				f.Printf("Divisível por %d na posição %d\n", v2[j], j)
				encontrou = true
			}
		}

		if !encontrou {
			f.Print("Não possui divisores no segundo vetor\n")
		}
	}
}
