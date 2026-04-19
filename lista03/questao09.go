package main

import f "fmt"

func main() {
	var (
		N                                 int
		quant_reprovados                  = 0
		quant_exames                      = 0
		quant_aprovados                   = 0
		n1, n2, media_aluno, media_classe float64
	)

	for {
		f.Print("\nDigite a quantidade de alunos a serem analisados: ")
		f.Scan(&N)
		if N <= 0 {
			f.Print("\nNúmero inválido! Digite um valor maior que zero.")
			continue
		} else {
			break
		}
	}

	medias := make([]float64, N)
	mensagens := make([]string, N)

	for i := 0; i < N; i++ {
		f.Print("\n")
		for {
			f.Printf("Digite a 1ª nota do aluno nº %d: ", i+1)
			f.Scan(&n1)
			if n1 < 0 || n1 > 10 {
				f.Print("\nNota inválida! Digite um valor entre 0 e 10.\n")
				continue
			} else {
				break
			}
		}

		for {
			f.Printf("Digite a 2ª nota do aluno nº %d: ", i+1)
			f.Scan(&n2)
			if n2 < 0 || n2 > 10 {
				f.Print("\nNota inválida! Digite um valor entre 0 e 10.\n")
				continue
			} else {
				break
			}
		}

		media_aluno = (n1 + n2) / 2
		medias[i] = media_aluno

		if media_aluno < 3 {
			mensagens[i] = "Reprovado"
			quant_reprovados++
		} else if media_aluno < 7 {
			mensagens[i] = "Exame"
			quant_exames++
		} else {
			mensagens[i] = "Aprovado"
			quant_aprovados++
		}
	}

	soma_medias := 0.0
	f.Print("\n----------- VISÃO GERAL  ------------")
	for j := 0; j < N; j++ {
		soma_medias += medias[j]
		f.Printf("\n	A média do %dº aluno é: %.2f | %v", j+1, medias[j], mensagens[j])
	}

	media_classe = soma_medias / float64(N)
	f.Print("\n--------------------------------------")
	f.Printf("\n	Total de alunos Aprovados: %d", quant_aprovados)
	f.Printf("\n	Total de alunos de Exame: %d", quant_exames)
	f.Printf("\n	Total de alunos Reprovados: %d", quant_reprovados)
	f.Printf("\n	Média da Classe: %.2f\n\n", media_classe)
}
