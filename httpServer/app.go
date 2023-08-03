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

func headers(w http.ResponseWriter, req *http.Request){
	log.Println("Requisição recebida em /headers")
	for name, headers := range req.Header{
		for _, h := range headers{
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

/*
ROADMAP

1. conectar com o banco de dados

2. implementar a API
- Adicionar todo
- Remover todo
- Ver todas as todos
- Check em uma todo
*/