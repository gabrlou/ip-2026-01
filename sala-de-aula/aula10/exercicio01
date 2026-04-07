package main

import f "fmt"

func main() {

	type pessoa struct {
		Nome       string
		Altura     float64
		Peso_ideal float64
	}

	quant := 0
	f.Print("Quantas pessoas serão cadastradas? ")
	f.Scanln(&quant)

	pessoas := make([]pessoa, quant)

	var dados pessoa

	for i := 0; i < len(pessoas); i++ {
		f.Printf("\nDigite o nome da %dª pessoa: ", i+1)
		f.Scanln(&dados.Nome)

		if dados.Nome == "FIM" {
			break
		}

		f.Printf("Digite a altura da %dª pessoa: ", i+1)
		f.Scanln(&dados.Altura)

		dados.Peso_ideal = 72.7*dados.Altura - 58

		f.Print("\n")

		pessoas[i] = dados
	}
	f.Print(pessoas)

	for j := 0; j <= len(pessoas); j++ {
		f.Printf("Pessoa: %d\nNome: %v\nAltura: %.2f\nPeso ideal: %.2f\n\n", j+1, pessoas[j])
	}
}
