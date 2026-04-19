package main

import (
	f "fmt"
)

func main() {
	var (
		preco        float64
		ingressos    int
		lucro        float64
		lucro_max    float64
		melhor_preco float64
		melhor_quant int
	)

	f.Printf("\nTabela de lucro esperado:\n\n")
	f.Printf("Preço\t\tIngressos\tLucro\n")

	preco = 6.0
	ingressos = 130

	for preco >= 1.0 {

		lucro = preco*float64(ingressos) - 300

		f.Printf("R$ %.2f\t\t%d\t\tR$ %.2f\n", preco, ingressos, lucro)

		if lucro > lucro_max || preco == 6.0 {
			lucro_max = lucro
			melhor_preco = preco
			melhor_quant = ingressos
		}

		preco -= 0.60
		ingressos += 30
	}

	f.Printf("\nLucro máximo esperado\n\n")
	f.Printf("	Lucro: R$ %.2f\n", lucro_max)
	f.Printf("	Preço do ingresso: R$ %.2f\n", melhor_preco)
	f.Printf("	Quantidade de ingressos: %d\n\n", melhor_quant)
}
