package main

import f "fmt"

func main() {
	vetor_1 := make([]int, 10)
	vetor_2 := mke([]int, 5)

	var resultante_1 []int
	var resultante_2 []int

	f.Print("\n")

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº número para ser adicionado ao primeiro vetor: ", i+1)
		f.Scan(&vetor_1[i])
	}

	f.Print("\n")

	for i := 0; i < 5; i++ {
		f.Printf("Digite o %dº número para ser adicionado ao segundo vetor: ", i+1)
		f.Scan(&vetor_2[i])
	}

	if vetor_1[i]%2 == 0 {
		maiorQue50[contador] = vetor[i]
		posicao[contador] = i + 1
		existe = true
		contador++
	}

	if existe {
		f.Print("\nNúmero(s) > 50\t\tPosição no Vetor")
		for i := 0; i < contador; i++ {
			f.Printf("\n%d\t\t\t%d", maiorQue50[i], posicao[i])
		}
		f.Print("\n")
	} else {
		f.Print("\nNão existem números maiores que 50 no vetor preenchido.\n\n")
	}
}
