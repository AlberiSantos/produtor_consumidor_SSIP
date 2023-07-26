# Hands-on

Tags: Feito

# Implementar uma aplicação que atenda ao seguinte caso de uso:

- Aplicação A usa Flask ou outro framework para expor um endpoint em uma dada porta, na VM criada na cloud
- Aplicação B envia uma requisição para o endpoint da aplicação A
- Aplicação B não deve estar na cloud, mas em uma máquina externa com acesso à rede da cloud
- Ambas as aplicações possuem um log capaz de identificar as operações que estão acontecendo
- Linguagem de programação de preferência pode ser utilizada

# Passo a Passo

## (Parte 3) Implementar uma aplicação que atenda ao caso de uso

## Preparando o ambiente

### Baixando e Instalando o Go

Para realização desta parte da atividade, foi escolhida a linguagem Go. Sendo assim, foi necessário instalar o Go na VM da cloud, segue abaixo os passos para isso:

1. Baixar o arquivo de instalação do Go diretamente na VM.
    
    ```bash
    wget https://go.dev/dl/go1.20.6.linux-amd64.tar.gz
    ```
    
    Neste caso estamos utilizando a versão 1.20.6. Que é a mais atual até o momento.
    
2. Descompactar o arquivo tar.gz baixado usando o comando `tar` .
    
    ```bash
    sudo tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz
    ```
    
3. Adicionar /usr/local/go/bin à `PATH` variável de ambiente.
    
    Isso pode ser feito adicionando a seguinte linha ao $HOME/.profile ou /etc/profile (para uma instalação em todo o sistema):
    
    ```bash
    export PATH=$PATH:/usr/local/go/bin
    ```
    
    **Observação:** as alterações feitas em um arquivo de perfil podem não ser aplicadas até a próxima vez que você fizer login no computador. Para aplicar as alterações imediatamente, basta executar os comandos do shell diretamente ou executá-los no perfil usando um comando como `source $HOME/.profile`.
    
4. Verificar se o Go foi instalado corretamente abrindo um prompt de comando e digitando o seguinte comando:
    
    ```bash
    go version
    ```
    

## Implementação da Aplicação A

Para implementar a aplicação responsável por expor um endpoint em uma dada porta, foi utilizado como referência o código disponível na documentação do Go, mais especificamente, no tutorial “[Go by Example: HTTP Server](https://gobyexample.com/http-server)”. 

Esse código em Go é um exemplo simples de um servidor HTTP que responde a duas rotas: "/hello" e "/headers". Aqui está uma explicação linha por linha do código:

```bash
package main

import (
        "fmt"
        "net/http"
)
```

- A declaração **`package main`** define que este arquivo faz parte do pacote principal do programa.
- As importações **`fmt`** e **`net/http`** são necessárias para utilizar funcionalidades relacionadas à formatação de saída e ao servidor HTTP, respectivamente.

```bash
func hello(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "hello\n")
}
```

- A função **`hello`** é um handler que será chamado quando a rota "/hello" for acessada. Ela recebe dois argumentos: **`w`** (http.ResponseWriter) e **`req`** (*http.Request). O **`w`** é usado para escrever a resposta HTTP, e o **`req`** contém informações da requisição recebida.
- Nesse caso, a função escreve "hello" na resposta usando **`fmt.Fprintf`**.

```bash
func headers(w http.ResponseWriter, req *http.Request){
        for name, headers := range req.Header{
                for _, h := range headers{
                        fmt.Fprintf(w, "%v: %v\n", name, h)
                }
        }
}
```

- A função **`headers`** é outro handler que será chamado quando a rota "/headers" for acessada. Ela também recebe os mesmos argumentos **`w`** e **`req`**.
- Essa função itera pelos headers da requisição **`req.Header`** usando um loop **`for range`**. Em seguida, outro loop **`for`** é usado para iterar pelos valores de cada header.
- A função escreve os nomes dos headers e seus valores na resposta.

```bash
func main() {
        http.HandleFunc("/hello", hello)
        http.HandleFunc("/headers", headers)

        http.ListenAndServe(":8090", nil)
}
```

- A função **`main`** é o ponto de entrada do programa.
- Nela, as funções **`http.HandleFunc`** são usadas para associar os handlers (**`hello`** e **`headers`**) às rotas ("/hello" e "/headers", respectivamente).
- Por fim, o servidor HTTP é iniciado com **`http.ListenAndServe`**, especificando a porta **`:8090`** em que o servidor deve escutar as requisições. O último argumento é definido como **`nil`**, o que significa que o servidor usará o multiplexador padrão (DefaultServeMux).

### Adicionando logging à aplicação

```bash
package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println("Requisição recebida em /hello")
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	log.Println("Requisição recebida em /headers")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	log.Println("Servidor iniciado na porta 8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

Foram adicionadas chamadas à função **`log.Println`** para registrar informações relevantes sobre as requisições recebidas em cada handler. O pacote **`log`** escreverá essas informações no console ou no arquivo de log, dependendo da configuração do ambiente.

Os registros de log incluem uma mensagem informando qual rota foi acessada. Essas mensagens serão úteis para identificar as operações que estão acontecendo em cada endpoint.

Além disso, também adicionamos um log ao iniciar o servidor para indicar que ele está funcionando e qual porta está sendo utilizada.

## Implementação da Aplicação B

Nessa implementação foi criada uma segunda aplicação em uma máquina externa com acesso à rede da cloud que faz uma requisição para o endpoint da Aplicação A. O código da aplicação:

```bash
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("http://10.11.19.159:8090/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Resposta do servidor:", string(body))
}
```

Este exemplo básico realiza uma requisição GET para o endpoint `/hello` da aplicação A e imprime a resposta recebida no console.

