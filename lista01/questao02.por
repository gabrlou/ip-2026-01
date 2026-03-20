programa {
  inclua biblioteca Matematica --> mat
  funcao inicio()
  {
     //Entrada
    real valor_popular, valor_geral, valor_arquibancada, valor_cadeiras,
    porcent_popular, porcent_geral, porcent_arquibancada, porcent_cadeiras, renda
    inteiro  total_ingressos, quant_jogos

    valor_popular= 1.00
    valor_geral= 5.00
    valor_arquibancada= 10.00
    valor_cadeiras= 20.00

    //Quantidades por categoria
    escreva("Quantos jogos serão analisados? ")
    leia(quant_jogos)
    escreva("\n")

    //Lista de renda
    real lista_renda[quant_jogos]
    inteiro i

    para(i=0; i < quant_jogos; i++)
    {

    escreva("Quantas pessoas compraram ingressos para o jogo nº ", i+1,"? ")
    leia(total_ingressos)

    escreva("Qual a porcentagem de ingressos vendidos na categoria popular no ", i+1,"º jogo? ")
    leia(porcent_popular)

    escreva("Qual a porcentagem de ingressos vendidos na categoria geral no ", i+1,"º jogo? ")
    leia(porcent_geral)

    escreva("Qual a porcentagem de ingressos vendidos na categoria arquibancada no ", i+1,"º jogo? ")
    leia(porcent_arquibancada)

    escreva("Qual a porcentagem de ingressos vendidos na categoria cadeiras no ", i+1,"º jogo? ")
    leia(porcent_cadeiras)


    //Cálculo de renda total
    renda= (((porcent_popular/100)*total_ingressos)*valor_popular) + (((porcent_geral/100)*total_ingressos)*valor_geral) + (((porcent_arquibancada/100)*total_ingressos)*valor_arquibancada) + (((porcent_cadeiras/100)*total_ingressos)*valor_cadeiras)
    lista_renda[i]= renda
 
    escreva("\n")
    }


    //Saída
    para(i=0; i< quant_jogos; i++)
    {
      escreva("A RENDA DO JOGO N. ", i+1, " E = ", mat.arredondar(lista_renda[i], 2),"\n")
    }
   

  }
}
