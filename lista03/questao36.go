package main

import f "fmt"

func main() {
	var num int

	for {
		f.Print("\nDigite um numero inteiro positivo: ")
		f.Scan(&num)

		if num < 0 {
			f.Println("Número inválido. Digite um número maior que zero.")
			continue
		} else {
			break
		}
	}

	if num == 0 {
		f.Println("\n	Numero na base 16: 0")
		return
	}

	var restos []int

	for num > 0 {
		restos = append(restos, num%16)
		num = num / 16
	}

	f.Print("\n	Numero na base 16: ")
	for i := len(restos) - 1; i >= 0; i-- {
		f.Print(restos[i])
	}
	f.Print("\n\n")
}
