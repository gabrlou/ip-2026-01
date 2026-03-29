package main

import (
	f "fmt"
	m "math"
)

func main() {
	var altura, comprimento_aresta, volume, area_base float64

	//Entrada
	f.Print("Digite a altura da pirâmide: ")
	f.Scanln(&altura)

	f.Print("Digite o comprimento de uma aresta da pirâmide: ")
	f.Scanln(&comprimento_aresta)

	//Processamento
	area_base = (3 * m.Pow(comprimento_aresta, 2) * m.Sqrt(3)) / 2
	volume = (area_base * altura) / 3

	//Saída
	f.Print("\nO VOLUME DA PIRAMIDE E = ", f.Sprintf("%.2f", volume), " METROS CUBICOS\n")
}
