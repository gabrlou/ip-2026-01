package main

import f "fmt"

func main() {
	var (
		preco_normal, preco_final float64
		dia_semana, categoria     int
	)

	f.Print("Digite o preço normal de locação do DVD: ")
	f.Scanln(&preco_normal)

	f.Print("Digite o dia da semana (\"1\" para domingo,\"2\" para segunda, \"3\" para terça, \n\"4\" para quarta, \"5\" para quinta, \"6\"para sexta ou \"7\" para sábado): ")
	f.Scanln(&dia_semana)
	if dia_semana < 1 || dia_semana > 7 {
		f.Print("\nValores inválidos!\nDigite \"1\" para domingo,\"2\" para segunda, \"3\" para terça, \n\"4\" para quarta, \"5\" para quinta, \"6\"para sexta ou \"7\" para sábado.")
		return
	}

	f.Print("Digite qual a categoria do DVD (\"1\" para comum ou \"2\" para lançamento): ")
	f.Scanln(&categoria)
	if categoria != 1 && categoria != 2 {
		f.Print("\nValores inválidos!\nDigite \"1\" para comum ou \"2\" para lançamento.")
		return
	}

	if categoria == 1 { //Categoria comum
		if dia_semana == 4 || dia_semana == 6 || dia_semana == 7 || dia_semana == 1 { //Dia normal
			preco_final = preco_normal
			f.Printf("\nO preço de locação é de: R$ %.2f", preco_final)
		} else { //Dia com desconto
			preco_final = preco_normal - (preco_normal * 0.4)
			f.Printf("\nO preço de locação é de: R$ %.2f", preco_final)
		}
	} else { //Categoria de lançamento
		preco_final = preco_normal + (preco_normal * 0.15)
		if dia_semana == 4 || dia_semana == 6 || dia_semana == 7 || dia_semana == 1 { //Dia normal
			f.Printf("\nO preço de locação é de: R$ %.2f", preco_final)
		} else { //Dia com desconto
			preco_final -= (preco_final * 0.4)
			f.Printf("\nO preço de locação é de: R$ %.2f", preco_final)
		}
	}
}
