package main

import f "fmt"

func main() {
	//Entrada
	var (
		valor_popular, valor_geral, valor_arquibancada, valor_cadeiras                float64
		porcent_popular, porcent_geral, porcent_arquibancada, porcent_cadeiras, renda float64
	)

	var total_ingressos, quant_jogos int

	valor_popular = 1.00
	valor_geral = 5.00
	valor_arquibancada = 10.00
	valor_cadeiras = 20.00

	//Quantidades por categoria
	f.Print("Quantos jogos serão analisados? ")
	f.Scanln(&quant_jogos)
	f.Println("\n")

	//Lista de renda
	lista_renda := make([]float64, quant_jogos)
	var i int

	for i = 0; i < quant_jogos; i++ {

		f.Print("Quantas pessoas compraram ingressos para o jogo nº ", i+1, "? ")
		f.Scan(&total_ingressos)

		f.Print("Qual a porcentagem de ingressos vendidos na categoria popular no ", i+1, "º jogo? ")
		f.Scan(&porcent_popular)

		f.Print("Qual a porcentagem de ingressos vendidos na categoria geral no ", i+1, "º jogo? ")
		f.Scan(&porcent_geral)

		f.Print("Qual a porcentagem de ingressos vendidos na categoria arquibancada no ", i+1, "º jogo? ")
		f.Scan(&porcent_arquibancada)

		f.Print("Qual a porcentagem de ingressos vendidos na categoria cadeiras no ", i+1, "º jogo? ")
		f.Scan(&porcent_cadeiras)

		//Cálculo de renda total
		renda = (((porcent_popular / 100) * float64(total_ingressos)) * valor_popular) +
			(((porcent_geral / 100) * float64(total_ingressos)) * valor_geral) +
			(((porcent_arquibancada / 100) * float64(total_ingressos)) * valor_arquibancada) +
			(((porcent_cadeiras / 100) * float64(total_ingressos)) * valor_cadeiras)

		lista_renda[i] = renda

		f.Println("\n")
	}

	//Saída
	for i = 0; i < quant_jogos; i++ {
		f.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i+1, lista_renda[i], "\n")
	}
	f.Print("\n")
}
