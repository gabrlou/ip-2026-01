programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
  
    real temp_f, temp_c, chuva_pol, chuva_mm

  //Entrada de Dados
  escreva("Digite a temperatura em Farenheit a ser convertida para Celsius: ")
  leia(temp_f)

  escreva("Digite a quantidade de chuva em polegadas a ser convertida para milímetros: ")
  leia(chuva_pol)

  escreva("\n")

  //Processamento
  temp_c = (5 * temp_f - 160) / 9
  chuva_mm = (chuva_pol * 25.4)

  //Saída
  escreva("O VALOR EM CELSIUS = ", mat.arredondar(temp_c, 2), "\n")
  escreva("A QUANTIDADE DE CHUVA E = ", mat.arredondar(chuva_mm, 2), "\n")
    
  }
}
