programa {
  funcao inicio() {

  inteiro horas, minutos, segundos, conversao

  //Entrada
  escreva("Digite a quantidade de horas: ")
  leia(horas)

  escreva("Digite a quantidade de minutos: ")
  leia(minutos)

  escreva("Digite a quantidade de segundos: ")
  leia(segundos)

  //Processamento
  
  conversao = (horas * 3600) + (minutos * 60) + segundos

  //Saída
  escreva("\nO TEMPO EM SEGUNDOS E = ", conversao, "\n")
  }
}
