package main
import f "fmt"

func main() {
	var (
		preco_normal, preco_final float64
		tipo_pagamento int
	)

	f.Print("Digite o preço normal de venda: ")
	f.Scanln(&preco_normal)

	f.Print("Digite a forma de pagamento:\n1 - À vista, dinheiro ou cheque\n2 - À vista ou cartão de crédito\n3 - Em 2 vezes\n4 - Em 3 vezes\n")
	f.Scanln(&tipo_pagamento)

	if tipo_pagamento < 1 || tipo_pagamento > 4 {
		f.Print("\nDígito inválido!\n1 - À vista, dinheiro ou cheque\n2 - À vista ou cartão de crédito\n3 - Em 2 vezes\n4 - Em 3 vezes")
		return
	} else if tipo_pagamento == 1 { //Tipo 1
		preco_final = preco_normal - (preco_normal*0.1)
		f.Printf("\nValor a ser pago: R$ %.2f", preco_final)
	} else if tipo_pagamento == 2 { //Tipo 2
		preco_final = preco_normal - (preco_normal*0.05)
		f.Printf("\nValor a ser pago: R$ %.2f", preco_final)
	} else if tipo_pagamento == 3 { //Tipo 3
		preco_final = preco_normal
		f.Printf("\nValor a ser pago: R$ %.2f", preco_final)
	} else { //Tipo 4
		preco_final = preco_normal + (preco_normal*0.1)
		f.Printf("\nValor a ser pago: R$ %.2f", preco_final)
	}
}
