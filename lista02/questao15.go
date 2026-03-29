package main

import f "fmt"

func main() {
	var dia, mes, ano int

	f.Print("Digite o dia: ")
	f.Scan(&dia)

	f.Print("Digite o mês: ")
	f.Scan(&mes)

	f.Print("Digite o ano: ")
	f.Scan(&ano)

	if mes == 1 {
		f.Println("\n", dia, "de janeiro de", ano)
	} else if mes == 2 {
		f.Println("\n", dia, "de fevereiro de", ano)
	} else if mes == 3 {
		f.Println("\n", dia, "de março de", ano)
	} else if mes == 4 {
		f.Println("\n", dia, "de abril de", ano)
	} else if mes == 5 {
		f.Println("\n", dia, "de maio de", ano)
	} else if mes == 6 {
		f.Println("\n", dia, "de junho de", ano)
	} else if mes == 7 {
		f.Println("\n", dia, "de julho de", ano)
	} else if mes == 8 {
		f.Println("\n", dia, "de agosto de", ano)
	} else if mes == 9 {
		f.Println("\n", dia, "de setembro de", ano)
	} else if mes == 10 {
		f.Println("\n", dia, "de outubro de", ano)
	} else if mes == 11 {
		f.Println("\n", dia, "de novembro de", ano)
	} else if mes == 12 {
		f.Println("\n", dia, "de dezembro de", ano)
	}
}
