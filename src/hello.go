// Pacote main
package main

//import de pacotes da lib padrão
import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// declaração de constantes
const monitoramentos = 5
const delay = 5

/*
Variaveis
	- Escopo é a mesma coisa de qualquer linguagem de gente
	- Existem diversar formas de declarar variáveis:
		- var nome = "Douglas"
		- nome := "Douglas"
		- var nome string = "Douglas"
		- var nome string
		De todas as formas a variável é fortemente tipada de acordo com o tipo do conteúdo que recebe.
		Quando é declarada sem valor, assume o valor padrão, sendo 0 para números, "" para string e false para bool
		Para números, int recebem o tipo int64, float recebem float64

Loops

	- Não existe while
	- Existe o for:
		- O for pode ser usado igual qualquer linguagem de gente:
			- for i := 0; i < 10; i++ {
				fmt.Println(i)
			}

			- for {
				fmt.Println("Loop infinito")
			}

			- for _, log := range logs {
				fmt.Println(log)
			}

		- O for também pode ser usado para iterar sobre um array, slice, string, map ou channel
		- O for sem parâmetros é um loop infinito


If e else são iguais qualquer linguagem de gente

O switch é igual qualquer linguagem de gente mas não precisa de break

Funções
    - func nomeDaFuncao(parametro tipo) tipoRetorno {}
	- func nomeDaFuncao(parametro tipo) (tipoRetorno, tipoRetorno2) {}
	Funções podem retornar mais de um valor


Retorno de funções
	Quando uma função retorna mais de um valor, é possível atribuir os valores retornados a variáveis ou utilizar o _ para ignorar o valor

Arrays e Slices
	Quase nunca utilizar Arrays e sim Slices
	Arrays são estruturas de dados que possuem tamanho fixo e tipo de dado fixo
	Slices são estruturas de dados que possuem tamanho variável e tipo de dado fixo igual vector do C++

I/O
	Praticamente igual ao C
	Quando usando sem modificador, o tipo de dado é definido na variável
	- fmt.Println("Hello, World!")
	- fmt.Printf("Hello, %s!\n", nome)
	- fmt.Scanf("%d", &comando)
	- fmt.Scan(&comando)

	Arquivos precisam de autorizações para criar, ler e escrever
	- arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	- leitor := bufio.NewReader(arquivo) (bufio é um pacote da lib padrão que facilita a leitura de arquivos)
	- arquivo.WriteString(log)
	- arquivo.Close()

*/

func main() {
	exibeInstroducao()
	for {
		exbibeMenu()
		comando := leComando()
		switch comando {
		case 1:

			iniciarMonitoramento()

		case 2:
			mostrarLogs(lerLogs())

		case 0:
			fmt.Println("Saindo do programa ...")
			os.Exit(0)

		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeInstroducao() {
	nome := "José"
	versao := 1.1
	fmt.Println("Hello, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comand escolhido foi", comandoLido)
	return comandoLido
}

func exbibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando ...")
	sites := leSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {

		for _, site := range sites {
			fmt.Println("Testando site", site)
			testeSite(site)
		}

		time.Sleep(delay * time.Second)
	}

}

func testeSite(site string) {
	resp, err := http.Get(site)
	var log string
	if err != nil {
		log = time.Now().Format(time.RFC3339) + " - Ocorreu um erro na requisicao " + site + "\n"
	}
	if resp.StatusCode == 200 {
		log = time.Now().Format(time.RFC3339) + " - Site " + site + " foi carregado com sucesso! Status Code: " + strconv.Itoa(resp.StatusCode) + "\n"
	} else {
		log = time.Now().Format(time.RFC3339) + " - Site " + site + " esta com problemas. Status Code: " + strconv.Itoa(resp.StatusCode) + "\n"
	}

	salvaLogs(log)
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo", err)
		return sites
	}
	leitor := bufio.NewReader(arquivo)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		if err == io.EOF {
			arquivo.Close()
			return sites
		}
		if err != nil {
			fmt.Println("Ocorreu um erro", err)
			continue
		}

		sites = append(sites, linha)
	}

}

func salvaLogs(log string) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo", err)
	} else {

		arquivo.WriteString(log)
		arquivo.Close()
	}

}

func lerLogs() []string {
	var logs []string
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo", err)
		return logs
	} else {
		leitor := bufio.NewReader(arquivo)

		for {
			linha, err := leitor.ReadString('\n')
			linha = strings.TrimSpace(linha)
			if err == io.EOF {
				arquivo.Close()
				return logs
			}
			if err != nil {
				fmt.Println("Ocorreu um erro", err)
				continue
			}

			logs = append(logs, linha)
		}
	}
}

func mostrarLogs(logs []string) {
	for _, log := range logs {
		fmt.Println(log)
	}
}
