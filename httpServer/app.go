package main

import (
	"src/dbConnect"
)

func main() {
	dbConnect.DeleteTodo(1)
	dbConnect.GetAllTodos()
	/*http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	log.Println("Servidor iniciado na porta 8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}*/
}

/*func hello(w http.ResponseWriter, req *http.Request) {
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
*/
