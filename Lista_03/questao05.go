package main

import f "fmt"

func main() {
	var (
		idade, num     int
		altura, peso   float64
		total_pessoas  = 0
		idade_maior_50 = 0
		idade_10_a_20  = 0
		soma_altura    = 0.0
		peso_menor_40  = 0
	)

	f.Print("\n| PROGRAMA INICIADO |\n")

	for {
		for {
			f.Print("\nDigite a idade da pessoa: ")
			f.Scan(&idade)
			if idade <= 0 {
				f.Print("\nIdade inválida! Digite uma idade maior que zero.\n")
				continue
			} else {
				break
			}
		}

		for {
			f.Print("Digite a altura da pessoa (em metros): ")
			f.Scan(&altura)
			if altura <= 0 {
				f.Print("\nAltura inválida! Digite uma altura maior que zero.\n")
				continue
			} else {
				break
			}
		}

		for {
			f.Print("Digite o peso da pessoa (em quilogramas): ")
			f.Scan(&peso)
			if peso <= 0 {
				f.Print("\nPeso inválido! Digite um peso maior que zero.\n")
				continue
			} else {
				break
			}
		}
		total_pessoas++

		if idade > 50 {
			idade_maior_50++
		}

		if idade >= 10 && idade <= 20 {
			idade_10_a_20++
			soma_altura += altura
		}

		if peso < 40 {
			peso_menor_40++
		}

		f.Print("\nDigite 1 para \"SIM\" e qualquer outro número para \"NÃO\".\nDeseja continuar digitando? ")
		f.Scan(&num)

		if num == 1 {
			continue
		} else {
			break
		}
	}

	media_alturas := soma_altura / float64(idade_10_a_20)
	porcentagem := (float64(peso_menor_40) / float64(total_pessoas)) * 100
	if peso_menor_40 == 0 {
		porcentagem = 0
	} else if idade_10_a_20 == 0 {
		media_alturas = 0
	}

	f.Printf("\n	Total de pessoas analisadas: %d", total_pessoas)
	f.Printf("\n	Quantidade de pessoas com idade superior a 50 anos: %d", idade_maior_50)
	f.Printf("\n	Média das alturas das pessoas com idade entre 10 e 20 anos: %.1f", media_alturas)
	f.Printf("\n	Porcentagem de pessoas com peso inferior a 40 quilos: %.2f%%", porcentagem)
	f.Print("\n\n| FIM DO PROGRAMA |\n\n")
}
