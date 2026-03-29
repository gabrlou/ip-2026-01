package main

import f "fmt"

func main() {
	var conta_cliente int
	var consumo_agua, valor_conta float64
	var tipo_consumidor string

	//Entrada
	f.Print("Digite o número da conta de água: ")
	f.Scanln(&conta_cliente)

	f.Print("Qual foi o consumo de água (em metros cúbicos)? ")
	f.Scanln(&consumo_agua)

	f.Print("Digite qual o tipo de consumidor (\"C\" - COMERCIAL, \"I\" -INDUSTRIAL ou \"R\" - RESIDENCIAL): ")
	f.Scanln(&tipo_consumidor)
	f.Print("\n")

	//Processamento e Saída
	if tipo_consumidor == "R" {
		valor_conta = 5 + (0.05 * consumo_agua)
		f.Print("CONTA = ", conta_cliente, "\n")
		f.Print("VALOR DA CONTA = ", f.Sprintf("%.2f", valor_conta), "\n")
	} else if tipo_consumidor == "C" && consumo_agua <= 80 {
		valor_conta = 500
		f.Print("CONTA = ", conta_cliente, "\n")
		f.Print("VALOR DA CONTA = ", f.Sprintf("%.2f", valor_conta), "\n")
	} else if tipo_consumidor == "C" && consumo_agua > 80 {
		valor_conta = 500 + (0.25 * (consumo_agua - 80))
		f.Print("CONTA = ", conta_cliente, "\n")
		f.Print("VALOR DA CONTA = ", f.Sprintf("%.2f", valor_conta), "\n")
	} else if tipo_consumidor == "I" && consumo_agua <= 100 {
		valor_conta = 800
		f.Print("CONTA = ", conta_cliente, "\n")
		f.Print("VALOR DA CONTA = ", f.Sprintf("%.2f", valor_conta), "\n")
	} else if tipo_consumidor == "I" && consumo_agua > 100 {
		valor_conta = 800 + (0.04 * (consumo_agua - 100))
		f.Print("CONTA = ", conta_cliente, "\n")
		f.Print("VALOR DA CONTA = ", f.Sprintf("%.2f", valor_conta), "\n")
	}
}
