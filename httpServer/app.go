package main

import (
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/add", hello)
	//http.HandleFunc("/delete", headers)
	//http.HandleFunc("/check", headers)
	//http.HandleFunc("/get-all", headers)

	log.Println("Servidor iniciado na porta 8090")
	err := http.ListenAndServeTLS(":8090", ".secrets/server.crt", ".secrets/server.key", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
