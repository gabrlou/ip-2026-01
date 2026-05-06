package main

import f "fmt"

func main() {

	var codigo int
	vetor := make([]float64, 10)

	f.Print("\nDigite o código (0, 1 ou 2): ")
	f.Scan(&codigo)

	if codigo == 0 {
		f.Print("Programa encerrado.\n")
		return
	}

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº número: ", i+1)
		f.Scan(&vetor[i])
	}

	if codigo == 1 {
		f.Print("\nVetor na ordem direta:\n")
		for i := 0; i < 10; i++ {
			f.Printf("%.2f ", vetor[i])
		}

	} else if codigo == 2 {
		f.Print("\nVetor na ordem inversa:\n")
		for i := 9; i >= 0; i-- {
			f.Printf("%.2f ", vetor[i])
		}

	} else {
		f.Print("\nCódigo inválido.\n")
	}

	f.Print("\n\n")
}
