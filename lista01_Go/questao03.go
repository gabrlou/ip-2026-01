package main

import f "fmt"

func main() {

	var n1, n2, n3, numero int

	//Entradas e Validação
	f.Println("Digite o primeiro número (somente 1 dígito): ")
	f.Scan(&n1)

	if n1 > 9 || n1 < 0 {
		f.Println("DÍGITO INVÁLIDO")
		return
	}

	f.Println("Digite o segundo número (somente 1 dígito): ")
	f.Scan(&n2)

	if n2 > 9 || n2 < 0 {
		f.Println("DÍGITO INVÁLIDO")
		return
	}

	f.Println("Digite o terceiro número (somente 1 dígito): ")
	f.Scan(&n3)

	if n3 > 9 || n3 < 0 {
		f.Println("DÍGITO INVÁLIDO")
		return
	}

	//Processamento
	numero = n1*100 + n2*10 + n3

	//Saída
	f.Print("\n")
	f.Println(numero, ", ", numero*numero)
}
