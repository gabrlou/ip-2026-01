programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {

  real salario, salario_reajustado

  //Entrada
  escreva("Digite o salário do funcionário: ")
  leia(salario)

  se (salario <= 300) {
    salario_reajustado = salario + (salario * 0.5)
    escreva("\nSALARIO COM REAJUSTE = ", mat.arredondar(salario_reajustado,2), "\n")
  }

  senao {
    salario_reajustado = salario + (salario * 0.3)
    escreva("\nSALARIO COM REAJUSTE = ", mat.arredondar(salario_reajustado,2), "\n")
  }
  }
}
