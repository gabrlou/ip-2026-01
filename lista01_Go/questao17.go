package main

import f "fmt"

func main() {
	var n1, n2, contador int

	//Entrada
	f.Print("Digite um primeiro número inteiro: ")
	f.Scanln(&n1)

	f.Print("Digite um segundo número inteiro: ")
	f.Scanln(&n2)

	f.Print("\n")
	contador = 1

	if n1%2 == 0 {
		for contador <= n2 {
			f.Print(n1, " ")
			n1 += 2
			contador += 1
		}
		f.Print("\n")
	} else {
		f.Print("O PRIMEIRO NUMERO NAO E PAR\n")
	}
}
