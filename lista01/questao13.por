programa {
  funcao inicio() {

  real nota
  cadeia conceito

  //Entrada
  escreva("Digite a nota a ser convertida para conceito: ")
  leia(nota)

  //Processamento e Saída
  se (nota >= 0 e nota < 6) {
    escreva("\nNOTA = ", nota,"   CONCEITO = D\n")
  }

  senao se (nota >= 6 e nota < 7.5) {
    escreva("\nNOTA = ", nota,"   CONCEITO = C\n")
  }

  senao se (nota >= 7.5 e nota < 9) {
    escreva("\nNOTA = ", nota,"   CONCEITO = B\n")
  }

  senao {
    escreva("\nNOTA = ", nota,"   CONCEITO = A\n")
  }
  }
}
