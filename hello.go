package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {

	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://google.com.br"
	sites[2] = "https://www.alura.com.br"
	fmt.Println(sites)
	fmt.Println(reflect.TypeOf(sites))
	fmt.Println(sites)
	exibeNomes()
	//exibeIntroducao()
	for {
		//exibirMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
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

	//if comando == 1 {
	//    fmt.Println("Monitorando...")
	//} else if comando == 2 {
	//    fmt.Println("Exibindo Logs...")
	//} else if comando == 0 {
	//    fmt.Println("Saindo do programa...")
	//}else{
	//	fmt.Println("Não conheço esta comando")
	//}

}

func exibeIntroducao() (string, int) {
	nome := "Gabriel"
	idade := 25
	versao := 1.26
	fmt.Println("Olá, sr. " + nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("Olá, sr", nome, "sua idade é", idade, "anos")
	return nome, idade
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

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://google.com.br"
	sites[2] = "https://www.alura.com.br"

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code:",
			resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"Gabriel", "Daniel", "Maria"}
	fmt.Print(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida")

	fmt.Print(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
