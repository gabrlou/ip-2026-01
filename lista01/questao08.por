programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
  
    real custo, raio, altura, m_quadrados, pi, At, Ac, Al
    pi = 3.14159

  //Entrada de Dados
  escreva("Digite o raio da lata (em metros): ")
  leia(raio)

  escreva("Digite a altura da lata (em metros): ")
  leia(altura)

  escreva("\n")

  //Processamento
  Ac = pi * (raio * raio)
  Al = 2 * pi * raio * altura
  At = (2 * Ac) + Al
  
  custo = At * 100

  //Saída
  escreva("O VALOR DO CUSTO E = ", mat.arredondar(custo, 2), "\n")
  }
}
