package main

import f "fmt"

func main() {
	var n1, n2, n3, media float64

	f.Println("Digite a primeira nota: ")
	f.Scan(&n1)

	f.Println("Digite a primeira nota: ")
	f.Scan(&n2)

	f.Println("Digite a primeira nota: ")
	f.Scan(&n3)
	f.Println("\n")

	//Processamento
	media = (n1 + n2 + n3) / 3

	//Saída
	if media < 6 {
		f.Println("MÉDIA = ", f.Sprintf("%.2f", media), "\nREPROVADO\n")
	} else {
		f.Println("MÉDIA = ", f.Sprintf("%.2f", media), "\nAPROVADO\n")
	}

}
