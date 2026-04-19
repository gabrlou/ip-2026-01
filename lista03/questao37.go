package main

import f "fmt"

func main() {
	var num int

	for {
		f.Print("\nDigite um numero inteiro positivo na base 8: ")
		f.Scan(&num)

		if num < 0 {
			f.Println("Número inválido. Digite um número maior que zero.\n")
			continue
		} else {
			break
		}
	}

	resultado := 0
	potencia := 1 // 8^0

	for num > 0 {
		digito := num % 10
		resultado += digito * potencia
		potencia *= 8
		num = num / 10
	}

	f.Printf("\n	Numero na base 10: %d\n\n", resultado)
}
