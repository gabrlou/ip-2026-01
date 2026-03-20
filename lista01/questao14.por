programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {

  real altura, comprimento_aresta, volume, area_base

  //Entrada
  escreva("Digite a altura da pirâmide: ")
  leia(altura)

  escreva("Digite o comprimento de uma aresta da pirâmide: ")
  leia(comprimento_aresta)

  //Processamento
  area_base = (3 * (comprimento_aresta*comprimento_aresta) * mat.raiz(3,2)) / 2
  volume = (area_base * altura) / 3

  //Saída
  escreva("\nO VOLUME DA PIRAMIDE E = ", mat.arredondar(volume,2), " METROS CUBICOS\n")
  }
}
