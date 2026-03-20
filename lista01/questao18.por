programa {
  funcao inicio() {

  inteiro valor_inicial, razao, n_elementos, soma_n, contador, termo

  //Entrada
  escreva("Digite o valor inicial da progressão aritmética: ")
  leia(valor_inicial)

  escreva("Digite a razão da progressão aritmética: ")
  leia(razao)

  escreva("Digite a quantidade \"n\" de elementos da progressão a serem somados: ")
  leia(n_elementos)

  termo = valor_inicial
  soma_n = 0
  contador = 1
  
  //Processamento
  enquanto(contador <= n_elementos) {
      soma_n += termo
      termo += razao
      contador += 1
    } 

  //Saída
  escreva("\n", soma_n, "\n")
  }
}
