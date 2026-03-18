programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
  
  inteiro conta_cliente
  real consumo_agua, valor_conta
  cadeia tipo_consumidor

  //Entrada
  escreva("Digite o número da conta de água: ")
  leia(conta_cliente)

  escreva("Qual foi o consumo de água (em metros cúbicos)? ")
  leia(consumo_agua)

  escreva("Digite qual o tipo de consumidor (\"C\" - COMERCIAL, \"I\" -INDUSTRIAL ou \"R\" - RESIDENCIAL): ")
  leia(tipo_consumidor)
  escreva("\n")
  
  //Processamento e Saída

  //Consumidor Residencial
  se (tipo_consumidor == "R") 
  {
    valor_conta = 5 + (0.05*consumo_agua)
    escreva("CONTA = ", conta_cliente, "\n")
    escreva("VALOR DA CONTA = ", mat.arredondar(valor_conta,2))
  }

  //Consumidor Comercial
  senao se (tipo_consumidor == "C" e consumo_agua <= 80) 
  {
    valor_conta = 500
    escreva("CONTA = ", conta_cliente, "\n")
    escreva("VALOR DA CONTA = ", mat.arredondar(valor_conta,2))
  }

  senao se (tipo_consumidor == "C" e consumo_agua > 80) 
  {
    valor_conta = 500 + (0.25*(consumo_agua - 80))
    escreva("CONTA = ", conta_cliente, "\n")
    escreva("VALOR DA CONTA = ", mat.arredondar(valor_conta,2))
  }

  //Consumidor Industrial
  senao se (tipo_consumidor == "I" e consumo_agua <= 100) 
  {
    valor_conta = 800
    escreva("CONTA = ", conta_cliente, "\n")
    escreva("VALOR DA CONTA = ", mat.arredondar(valor_conta,2))
  }

  senao se (tipo_consumidor == "I" e consumo_agua > 100) 
  {
    valor_conta = 800 + (0.04*(consumo_agua - 100))
    escreva("CONTA = ", conta_cliente, "\n")
    escreva("VALOR DA CONTA = ", mat.arredondar(valor_conta,2))
  }

  }
}
