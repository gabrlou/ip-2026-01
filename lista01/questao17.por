programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {

  inteiro n1, n2, contador

  //Entrada
  escreva("Digite um primeiro número inteiro: ")
  leia(n1)

  escreva("Digite um segundo número inteiro: ")
  leia(n2)

  escreva("\n")
  contador = 1

  se (n1 % 2 == 0) {
    enquanto(contador <= n2) {
      escreva(n1, " ")
      n1 += 2
      contador += 1
    } 
    escreva("\n")
  }

  senao {
    escreva("O PRIMEIRO NUMERO NAO E PAR\n")
  }
  }
}
