package main

import f "fmt"

func main() {
	var valor_inicial, razao, n_elementos, soma_n, contador, termo int

	//Entrada
	f.Print("Digite o valor inicial da progressão aritmética: ")
	f.Scanln(&valor_inicial)

	f.Print("Digite a razão da progressão aritmética: ")
	f.Scanln(&razao)

	f.Print("Digite a quantidade \"n\" de elementos da progressão a serem somados: ")
	f.Scanln(&n_elementos)

	termo = valor_inicial
	soma_n = 0
	contador = 1

	//Processamento
	for contador <= n_elementos {
		soma_n += termo
		termo += razao
		contador += 1
	}

	//Saída
	f.Print("\n", soma_n, "\n")
}
