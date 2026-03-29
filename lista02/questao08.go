package main

import f "fmt"

func main() {
	var numero int

	f.Print("Digite um número inteiro: ")
	f.Scan(&numero)
	f.Print("\n")

	if numero > 20 && numero < 90 {
		f.Println("O número está entre 20 e 90.")
	} else {
		f.Println("O número não está entre 20 e 90.")
	}
	f.Print("\n")
}
