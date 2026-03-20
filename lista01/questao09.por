programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {

  real A, B, C, delta

  //Entrada
  escreva("Digite o coeficiente \"A\": ")
  leia(A)

  escreva("Digite o coeficiente \"B\": ")
  leia(B)

  escreva("Digite o coeficiente \"C\": ")
  leia(C)

  //Cálculo de Delta
  delta = (B*B) - 4 * A * C

  //Saída
	escreva("\n")
  escreva("O VALOR DE DELTA E = ", mat.arredondar(delta,2))
  }
}
