package main

import f "fmt"

func calculaDigito1(cpf []int) int {
	soma := 0
	peso := 2

	for i := 8; i >= 0; i-- {
		soma += cpf[i] * peso
		peso++
	}

	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}

func calculaDigito2(cpf []int, d1 int) int {
	soma := 0
	peso := 3

	for i := 8; i >= 0; i-- {
		soma += cpf[i] * peso
		peso++
	}

	soma += d1 * 2

	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}

func main() {
	var entrada string
	cpf := make([]int, 11)

	for {
		f.Print("\nDigite o CPF (apenas números): ")
		f.Scan(&entrada)

		if len(entrada) != 11 {
			f.Print("CPF inválido (digite um CPF com 11 dígitos).\n")
			continue
		} else {
			break
		}
	}

	for i := 0; i < 11; i++ {
		cpf[i] = int(entrada[i] - '0')
	}

	d1 := calculaDigito1(cpf)
	d2 := calculaDigito2(cpf, d1)

	if d1 == cpf[9] && d2 == cpf[10] {
		f.Print("\n	CPF válido!\n\n")
	} else {
		f.Print("\n	CPF inválido!\n\n")
	}
}
