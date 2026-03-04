package main

import (
	"fmt"
)

func main() {
	var nome string = "Gabriel"
	var idade int = 26
	var versao float32 = 1.26
	fmt.Println("Olá, sr. " + nome)
	fmt.Println("Olá, sr", nome, "sua idade é", idade, "anos")

	fmt.Printf("O tipo de variavel versão é %T\n", versao)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O endereço da minha variavel comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)

}
