package main

import (
	"fmt"
)

func main() {
	nome := "Gabriel"
	idade := 26
	versao := 1.26

	fmt.Println("Olá, sr. " + nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("Olá, sr", nome, "sua idade é", idade, "anos")

	fmt.Printf("O tipo de variavel versão é %T\n", versao)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Não conheço esta comando")
	}

}
