programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    
    real salario_minimo, consumo_kw, valor_por_kw, consumo_valor, consumo_com_desconto

    //Entrada
    escreva("Qual o valor do salário mínimo? ")
    leia(salario_minimo)

    escreva("Quantos kW foram gastos? ")
    leia(consumo_kw)

    //Processanto
    valor_por_kw = (0.7 * salario_minimo)/100
    consumo_valor = consumo_kw * valor_por_kw
    consumo_com_desconto = consumo_valor - (consumo_valor * 0.1)
    escreva("\n")

    //Saída
    escreva("Custo por kW: R$ ", mat.arredondar(valor_por_kw,2),"\n")
    escreva("Custo do consumo: R$ ", mat.arredondar(consumo_valor,2),"\n")
    escreva("Custo com desconto: R$ ", mat.arredondar(consumo_com_desconto,2),"\n")
    
  }
}
