package main

import f "fmt"

func main() {
	var temp_f, temp_c, chuva_pol, chuva_mm float64

	//Entrada de Dados
	f.Print("Digite a temperatura em Farenheit a ser convertida para Celsius: ")
	f.Scanln(&temp_f)

	f.Print("Digite a quantidade de chuva em polegadas a ser convertida para milímetros: ")
	f.Scanln(&chuva_pol)

	f.Print("\n")

	//Processamento
	temp_c = (5*temp_f - 160) / 9
	chuva_mm = (chuva_pol * 25.4)

	//Saída
	f.Print("O VALOR EM CELSIUS = ", f.Sprintf("%.2f", temp_c), "\n")
	f.Print("A QUANTIDADE DE CHUVA E = ", f.Sprintf("%.2f", chuva_mm), "\n")
}
