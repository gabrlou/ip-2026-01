package main

import f "fmt"

func main() {
	var destino, idaVolta int

	f.Print("\nQual o seu destino?\nDigite \"1\" para a Região Norte, \"2\" para a Região Nordeste,\n\"3\" para a Região Centro-Oeste e \"4\" para a Região Sul: ")
	f.Scan(&destino)

	f.Print("\nA viagem inclui retorno?\nDigite \"1\" para sim (ida e volta) e \"2\" para não (só volta): ")
	f.Scan(&idaVolta)

	if destino != 1 && destino != 2 && destino != 3 && destino != 4 {
		f.Print("Destino inválido. Por favor, escolha um número entre 1 e 4.")
		return
	} else if idaVolta != 1 && idaVolta != 2 {
		f.Print("Opção de ida e volta inválida. Por favor, escolha \"1\" para sim ou \"2\" para não.")
		return
	} else if destino == 1 && idaVolta == 1 {
		f.Print("O valor da passagem para a Região Norte (ida e volta) é R$ 900,00.")
	} else if destino == 1 && idaVolta == 2 {
		f.Print("O valor da passagem para a Região Norte (só ida) é R$ 500,00.")
	} else if destino == 2 && idaVolta == 1 {
		f.Print("O valor da passagem para a Região Nordeste (ida e volta) é R$ 650,00.")
	} else if destino == 2 && idaVolta == 2 {
		f.Print("O valor da passagem para a Região Nordeste (só ida) é R$ 350,00.")
	} else if destino == 3 && idaVolta == 1 {
		f.Print("O valor da passagem para a Região Centro-Oeste (ida e volta) é R$ 600,00.")
	} else if destino == 3 && idaVolta == 2 {
		f.Print("O valor da passagem para a Região Centro-Oeste (só ida) é R$ 350,00.")
	} else if destino == 4 && idaVolta == 1 {
		f.Print("O valor da passagem para a Região Sul (ida e volta) é R$ 550,00.")
	} else if destino == 4 && idaVolta == 2 {
		f.Print("O valor da passagem para a Região Sul (só ida) é R$ 300,00.")
	}
}
