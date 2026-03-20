programa {
  funcao inicio() {

  inteiro horas, custo, blocos, resto
  //Entrada
  escreva("Digite a quantidade de horas de uso da charrete: ")
  leia(horas)

  //Processamento
  blocos = horas / 3
  resto = horas % 3
  custo = (10 * blocos) + (5 * resto)

  //Saída
  escreva("\nO VALOR A PAGAR E = ", custo,"\n")
  
  }
}
