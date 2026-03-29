package main

import f "fmt"

func main() {
	var horas, custo, blocos, resto int

	//Entrada
	f.Print("Digite a quantidade de horas de uso da charrete: ")
	f.Scanln(&horas)

	//Processamento
	blocos = horas / 3
	resto = horas % 3
	custo = (10 * blocos) + (5 * resto)

	//Saída
	f.Print("\nO VALOR A PAGAR E = ", custo, "\n")
}
