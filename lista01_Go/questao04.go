package main

import f "fmt"

func main() {
	var salario_minimo, consumo_kw, valor_por_kw, consumo_valor, consumo_com_desconto float64

	//Entrada
	f.Print("Qual o valor do salário mínimo? ")
	f.Scan(&salario_minimo)

	f.Print("Quantos kW foram gastos? ")
	f.Scan(&consumo_kw)

	//Processanto
	valor_por_kw = (0.7 * salario_minimo) / 100
	consumo_valor = consumo_kw * valor_por_kw
	consumo_com_desconto = consumo_valor - (consumo_valor * 0.1)
	f.Print("\n")

	//Saída
	f.Print("Custo por kW: R$ ", f.Sprintf("%.2f", valor_por_kw), "\n")
	f.Print("Custo do consumo: R$ ", f.Sprintf("%.2f", consumo_valor), "\n")
	f.Print("Custo com desconto: R$ ", f.Sprintf("%.2f", consumo_com_desconto), "\n")
}
