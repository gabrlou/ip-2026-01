package main

import f "fmt"

func main() {

	vetor := make([]int, 10)

	f.Print("\n")
	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº número: ", i+1)
		f.Scan(&vetor[i])
	}

	f.Print("\nNúmeros primos e suas posições:\n")

	for i := 0; i < 10; i++ {

		n := vetor[i]

		if n < 2 {
			continue
		}

		primo := true
		for j := 2; j*j <= n; j++ {
			if n%j == 0 {
				primo = false
				break
			}
		}

		if primo {
			f.Printf("Número %d na posição %d\n", n, i)
		}
	}

	f.Print("\n\n")
}
