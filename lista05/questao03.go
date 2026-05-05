package main

import f "fmt"

func main() {

	vetor := make([]int, 10)

	var pares []int
	var impares []int
	somaPares := 0

	f.Print("\n")
	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº número: ", i+1)
		f.Scan(&vetor[i])
	}

	for i := 0; i < 10; i++ {
		if vetor[i]%2 == 0 {
			pares = append(pares, vetor[i])
			somaPares += vetor[i]
		} else {
			impares = append(impares, vetor[i])
		}
	}

	f.Print("\n	Números pares digitados: ", pares)
	f.Print("\n	Soma dos pares: ", somaPares)
	f.Print("\n	Números ímpares digitados: ", impares)
	f.Print("\n	Quantidade de ímpares: ", len(impares))
	f.Print("\n\n")
}
