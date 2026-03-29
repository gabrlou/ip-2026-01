package main

import f "fmt"

func main() {
	var valor_compra, valor_venda float64

	f.Print("Digite o valor de compra: ")
	f.Scan(&valor_compra)

	if valor_compra <= 0 {
		f.Println("\nValor Inválido!\nO valor de compra deve ser maior que zero.")
		return
	} else if valor_compra < 10 {
		valor_venda = (valor_compra * 0.7) + valor_compra
		f.Printf("\nValor de venda: %.2f\n", valor_venda)
	} else if valor_compra >= 10 && valor_compra < 30 {
		valor_venda = (valor_compra * 0.5) + valor_compra
		f.Printf("\nValor de venda: %.2f\n", valor_venda)
	} else if valor_compra >= 30 && valor_compra < 50 {
		valor_venda = (valor_compra * 0.4) + valor_compra
		f.Printf("\nValor de venda: %.2f\n", valor_venda)
	} else {
		valor_venda = (valor_compra * 0.3) + valor_compra
		f.Printf("\nValor de venda: %.2f\n", valor_venda)
	}
}
