programa {
    inclua biblioteca Matematica --> mat
  funcao inicio()
  {
     //Entrada
    real n1, n2 , n3, total, media, media_2_casas
    escreva("Digite a primeira nota: ")
    leia(n1)
    
    escreva("Digite a segunda nota: ")
    leia(n2)
    
    escreva("Digite a terceira nota: ")
    leia(n3)
    escreva("\n")


    //Processamento
    total= n1 + n2 + n3
    media= total / 3
    media_2_casas= mat.arredondar(media, 2)

    //Saída
    se (media_2_casas < 6)
    {
      escreva("\n" , "MÉDIA = ", media_2_casas,"\n", "REPROVADO")
    }
 
    senao
   {
    escreva("MÉDIA = ", media_2_casas,"\n", "APROVADO", "\n")
   }

  }
}
