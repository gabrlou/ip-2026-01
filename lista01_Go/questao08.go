package main

import f "fmt"

func main() {
	var custo, raio, altura, pi, At, Ac, Al float64
	pi = 3.14159

	//Entrada de Dados
	f.Print("Digite o raio da lata (em metros): ")
	f.Scanln(&raio)

	f.Print("Digite a altura da lata (em metros): ")
	f.Scanln(&altura)

	f.Print("\n")

	//Processamento
	Ac = pi * (raio * raio)
	Al = 2 * pi * raio * altura
	At = (2 * Ac) + Al

	custo = At * 100

	//Saída
	f.Print("O VALOR DO CUSTO E = ", f.Sprintf("%.2f", custo), "\n")
}
