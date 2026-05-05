package main

import f "fmt"

func main() {

	vetor_1 := make([]int, 10)
	vetor_2 := make([]int, 5)

	var resultante_1 []int
	var resultante_2 []int

	f.Print("\n")
	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº número para o primeiro vetor: ", i+1)
		f.Scan(&vetor_1[i])
	}

	f.Print("\n")

	for i := 0; i < 5; i++ {
		f.Printf("Digite o %dº número para o segundo vetor: ", i+1)
		f.Scan(&vetor_2[i])
	}

	soma_v2 := 0
	for i := 0; i < 5; i++ {
		soma_v2 += vetor_2[i]
	}

	for i := 0; i < 10; i++ {
		if vetor_1[i]%2 == 0 {
			resultante_1 = append(resultante_1, vetor_1[i]+soma_v2)
		} else {
			resultante_2 = append(resultante_2, vetor_1[i]+soma_v2)
		}
	}

	f.Print("\nPrimeiro vetor resultante (pares): ", resultante_1)
	f.Print("\nSegundo vetor resultante (ímpares): ", resultante_2)
	f.Print("\n\n")
}
