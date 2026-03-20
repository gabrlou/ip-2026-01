programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {

  inteiro numero, i

  //Entrada
  escreva("Digite um número inteiro maior que 5 e menor que 2000: ")
  leia(numero)

  se (numero <= 5 ou numero >= 2000) {
    escreva("\nNÚMERO INVÁLIDO!\nDigite um número maior que 5 e menor que 2000.\n")
    retorne
  }

  escreva("\n")

  //Processamento e Saída
  para(i = 2; i <= numero; i += 2 ) {
    escreva(i,"^2 = ", i * i, "\n")
  }
  }
}
