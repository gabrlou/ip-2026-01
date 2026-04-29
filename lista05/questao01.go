package main

import f "fmt"

func main() {
	vetor := make([]int, 10)
	maiorQue50 := make([]int, 10)
	posicao := make([]int, 10)
	existe := false
	contador := 0

	f.Print("\n")

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº número para ser adicionado ao vetor: ", i+1)
		f.Scan(&vetor[i])

		if vetor[i] > 50 {
			maiorQue50[contador] = vetor[i]
			posicao[contador] = i + 1
			existe = true
			contador++
		}
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
