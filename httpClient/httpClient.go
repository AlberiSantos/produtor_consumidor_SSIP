package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
)

func main() {
        response, err := http.Get("http://10.11.19.156:8090/hello")
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

