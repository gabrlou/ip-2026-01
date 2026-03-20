programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {

  real a, b, c, d, determinante

  //Entrada
  escreva("Digite o elemente \"a\" da matriz 2x2: ")
  leia(a)

  escreva("Digite o elemento \"b\" da matriz 2x2: ")
  leia(b)

  escreva("Digite o elemento \"c\" da matriz 2x2: ")
  leia(c)

  escreva ("Digite o elemento \"d\" da matriz 2x2: ")
  leia(d)

  //Cálculo de Determinante
  determinante = (a * d) - (b * c)

  //Saída
  escreva("\n")
  escreva("O VALOR DO DETERMINANTE E = ", mat.arredondar(determinante,2))
  }
}
