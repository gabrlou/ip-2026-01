package main

import f "fmt"

func main() {
	var (
		num                 int
		primeiro            = true
		soma                = 0
		quant_num           = 0
		media               = 0.0
		maior_num           = 0
		menor_num           = 0
		quant_num_par       = 0
		soma_par            = 0
		media_pares         = 0.0
		quant_num_impar     = 0
		porcentagem_impares = 0.0
	)

	f.Print("\n| PROGRAMA INICIADO |")
	f.Print("\nPara finalizar, digite o número 30000. Caso contrário, digite qualquer outro número.\n")

	for {

		f.Print("\nDigite um número: ")
		f.Scan(&num)

		if num == 30000 {
			break
		}

		soma += num
		quant_num++

		if primeiro {
			maior_num = num
			menor_num = num
			primeiro = false
		} else {
			if num > maior_num {
				maior_num = num
			}

			if num < menor_num {
				menor_num = num
			}
		}

		if num%2 == 0 {
			quant_num_par++
			soma_par += num
		}

		if num%2 != 0 {
			quant_num_impar++
		}
	}

	if quant_num > 0 {
		media = float64(soma) / float64(quant_num)
	}

	if quant_num_par > 0 {
		media_pares = float64(soma_par) / float64(quant_num_par)
	}

	if quant_num_impar > 0 {
		porcentagem_impares = (float64(quant_num_impar) / float64(quant_num)) * 100
	}

	f.Printf("\n	Soma dos números digitados: %d", soma)
	f.Printf("\n	Total de números digitados: %d", quant_num)
	f.Printf("\n	Média dos números digitados: %.2f", media)
	f.Printf("\n	Maior número digitado: %d", maior_num)
	f.Printf("\n	Menor número digitado: %d", menor_num)
	f.Printf("\n	Média dos números pares digitados: %.2f", media_pares)
	f.Printf("\n	Porcentagem de números ímpares digitados: %.2f%%", porcentagem_impares)
	f.Print("\n\n| FIM DO PROGRAMA |\n\n")
}
