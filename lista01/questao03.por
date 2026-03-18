programa {
  funcao inicio() {
    
    inteiro n1, n2, n3, numero

    //Entradas e Validação
    escreva("Digite o primeiro número (somente 1 dígito): ")
    leia(n1)
    se (n1>9 ou n1<0)
    {
      escreva("DÍGITO INVÁLIDO")
      retorne
    }

    escreva("Digite o segundo número (somente 1 dígito): ")
    leia(n2)
    se (n2>9 ou n2<0)
    {
      escreva("DÍGITO INVÁLIDO")
      retorne
    }

    escreva("Digite o terceiro número (somente 1 dígito): ")
    leia(n3)
    se (n3>9 ou n3<0)
    {
      escreva("DÍGITO INVÁLIDO")
      retorne
    }

    //Processamento
    numero = n1*100 + n2*10 + n3

    //Saída
    escreva(numero, ", ", numero*numero)
    
  }
}
