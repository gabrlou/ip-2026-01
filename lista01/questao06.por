programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
  
    real temp_f, temp_c
    inteiro quant_convesoes

  //Quantidade de Conversões
  escreva("Digite a quantidade de conversões a serem executadas: ")
  leia(quant_convesoes)
  escreva("\n")

  real lista_far[quant_convesoes]
  inteiro i

  //Guardar as temperaturas numa lista
  para(i = 0; i <  quant_convesoes; i++)
    {
      escreva("Digite a ", i+1, "ª", " temperatura a ser convertida para Celsius: ")
      leia(temp_f)

      lista_far[i] = temp_f
    }

    escreva("\n")

  //Mostrar as conversões de cada temperatura para Celsius
  para(i = 0; i < quant_convesoes; i++)
    {
      temp_c = (5 * (lista_far[i] - 32)) / 9
    escreva(lista_far[i], " FAHRENHEIT EQUIVALE A ", mat.arredondar(temp_c,2), " CELSIUS \n")
    }

  }
}
