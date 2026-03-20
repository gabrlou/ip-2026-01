programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {

  inteiro numero, contador
  real soma

  //Entrada
  escreva("Digite um número inteiro, positivo e maior que 1: ")
  leia(numero)

  se (numero <= 0 ou numero == 1) {
    escreva("\nNumero invalido!\n")
    retorne
  }

  soma = 0
  contador = 1

  //Processamento
  enquanto(contador <= numero) {
      soma += 1 / contador
      contador += 1
    } 

  //Saída
  escreva("\n", mat.arredondar(soma,6), "\n")
  }
}
