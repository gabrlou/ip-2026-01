package main

import f "fmt"

func main() {
	var horas, minutos, segundos, conversao int

	//Entrada
	f.Print("Digite a quantidade de horas: ")
	f.Scanln(&horas)

	f.Print("Digite a quantidade de minutos: ")
	f.Scanln(&minutos)

	f.Print("Digite a quantidade de segundos: ")
	f.Scanln(&segundos)

	//Processamento

	conversao = (horas * 3600) + (minutos * 60) + segundos

	//Saída
	f.Print("\nO TEMPO EM SEGUNDOS E = ", conversao, "\n")
}
