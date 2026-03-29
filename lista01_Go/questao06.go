package main

import f "fmt"

func main() {
	var temp_f, temp_c float64
	var quant_convesoes int

	//Quantidade de Conversões
	f.Print("Digite a quantidade de conversões a serem executadas: ")
	f.Scanln(&quant_convesoes)
	f.Print("\n")

	lista_far := make([]float64, quant_convesoes)
	i := 0

	//Laço para guardar as temperaturas numa lista
	for i = 0; i < quant_convesoes; i++ {
		f.Print("Digite a ", i+1, "ª", " temperatura a ser convertida para Celsius: ")
		f.Scanln(&temp_f)

		lista_far[i] = temp_f
	}

	f.Print("\n")

	//Laço para mostrar as conversões de cada temperatura para Celsius
	for i = 0; i < quant_convesoes; i++ {
		temp_c = (5 * (lista_far[i] - 32)) / 9
		f.Print(lista_far[i], " FAHRENHEIT EQUIVALE A ", f.Sprintf("%.2f", temp_c), " CELSIUS \n")
	}
}
