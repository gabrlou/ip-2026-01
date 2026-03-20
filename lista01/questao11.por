programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {

  inteiro numero

  //Entrada
  escreva("Digite um número inteiro: ")
  leia(numero)

  se (numero % 3 == 0 e numero % 5 == 0) {
    escreva("\nO NUMERO E DIVISIVEL\n")
  }

  senao {
    escreva("\nO NUMERO NAO E DIVISIVEL\n")
  }
  }
}
