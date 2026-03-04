package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()
	leSitesDoArquivo()
	for {
		exibirMenu()

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
	fmt.Println("")

	return comandoLido
}

func exibirMenu() {
	fmt.Println("")
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitoramento...")

	//sites := []string{"https://random-status-code.herokuapp.com/",
	//	"https://google.com.br", "https://www.alura.com.br"}

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code:",
			resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)

	}

	leitor := bufio.NewReader(arquivo)

	linha, err := leitor.ReadString('\n')

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	fmt.Println(linha)

	return sites
}
