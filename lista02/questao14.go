package main

import f "fmt"

func main() {
	var preco_inicial, preco_final float64
	var ar, pint_met, vidro_elet, hidraulica int

	f.Print("Digite o preço inicial de fábrica do carro: ")
	f.Scanln(&preco_inicial)

	preco_final = preco_inicial

	f.Print("O carro possuirá Ar Condicionado? (1 para sim, 0 para não): ")
	f.Scanln(&ar)
	if ar != 1 && ar != 0 {
		f.Println("\nValor inválido para Ar Condicionado. Por favor, insira 1 para sim ou 0 para não.")
		return
	} else if ar == 1 {
		preco_final += 1750
	}

	f.Print("O carro possuirá Pintura Metálica? (1 para sim, 0 para não): ")
	f.Scanln(&pint_met)
	if pint_met != 1 && pint_met != 0 {
		f.Println("\nValor inválido para Pintura Metálica. Por favor, insira 1 para sim ou 0 para não.")
		return
	} else if pint_met == 1 {
		preco_final += 800
	}

	f.Print("O carro possuirá Vidros Elétricos? (1 para sim, 0 para não): ")
	f.Scanln(&vidro_elet)
	if vidro_elet != 1 && vidro_elet != 0 {
		f.Println("\nValor inválido para Vidros Elétricos. Por favor, insira 1 para sim ou 0 para não.")
		return
	} else if vidro_elet == 1 {
		preco_final += 1200
	}

	f.Print("O carro possuirá Direção Hidráulica? (1 para sim, 0 para não): ")
	f.Scanln(&hidraulica)
	if hidraulica != 1 && hidraulica != 0 {
		f.Println("\nValor inválido para Direção Hidráulica. Por favor, insira 1 para sim ou 0 para não.")
		return
	} else if hidraulica == 1 {
		preco_final += 2000
	}
	f.Println("\nO valor final do carro é de: R$ ", preco_final)
}
