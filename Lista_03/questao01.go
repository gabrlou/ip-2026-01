package main

import f "fmt"

func main() {
	var base, expoente, resultado int

	for {
		f.Print("\nDigite a base da potência: ")
		f.Scan(&base)

		if base < 0 {
			f.Print("\nValor inválido! Digite um número maior ou igual a zero.\n")
			continue
		}
		break
	}

	for {
		f.Print("Digite o expoente da potência: ")
		f.Scan(&expoente)

		if expoente < 0 {
			f.Print("\nValor inválido! Digite um número maior ou igual a zero.\n")
			continue
		}
		break
	}

	resultado = base
	for i := 1; i < expoente; i++ {
		resultado *= base
	}

	f.Printf("\nBase: %d | Expoente: %d\nResultado: %d\n\n", base, expoente, resultado)
}
