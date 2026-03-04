package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibirMenu()
	comando := leComando()

	fmt.Scanf("%d", &comando)

	//if comando == 1 {
	//    fmt.Println("Monitorando...")
	//} else if comando == 2 {
	//    fmt.Println("Exibindo Logs...")
	//} else if comando == 0 {
	//    fmt.Println("Saindo do programa...")
	//}else{
	//	fmt.Println("Não conheço esta comando")
	//}

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Print("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

}

func exibeIntroducao() {
	nome := "Gabriel"
	idade := 26
	versao := 1.26
	fmt.Println("Olá, sr. " + nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("Olá, sr", nome, "sua idade é", idade, "anos")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibirMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func saiDoPrograma() {

}
