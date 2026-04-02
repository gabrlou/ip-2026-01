package main

import f "fmt"

func main() {
	var idade int

	f.Print("Digite sua idade: ")
	f.Scanln(&idade)

	if idade < 0 {
		f.Print("\nIdade inválida!\nDigite uma idade maior que zero.")
	} else if idade < 16 {
		f.Print("\nClasse eleitoral: Não-eleitor")
	} else if (idade >= 16 && idade < 18) || idade > 65 {
		f.Print("\nClasse eleitoral: Eleitor facultativo")
	} else {
		f.Print("\nClasse eleitoral: Eleitor obrigatório")
	}
}
