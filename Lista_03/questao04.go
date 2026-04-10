package main

import f "fmt"

func main() {
	var (
		num       int
		resultado string
	)

	f.Print("\n| Programa Iniciado |\nPara finalizar, digite um número menor ou igual a 0\n")
	for {
		resultado = "falso"
		f.Print("\nDigite um número inteiro: ")
		f.Scan(&num)

		if num <= 0 {
			f.Print("\n| Programa Encerrado |\n\n")
			break
		} else {
			for i := 0; (i * i) <= num; i++ {
				if i*i == num {
					f.Printf("O número %d é um quadrado perfeito. | %d * %d = %d\n", num, i, i, num)
					resultado = "verdadeiro"
					break
				}
			}
		}
		if resultado == "falso" {
			f.Printf("O número %d não é um quadrado perfeito.\n", num)
		}
		continue
	}
}
