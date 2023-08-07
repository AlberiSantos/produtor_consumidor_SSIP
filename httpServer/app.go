package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"main/src/dbConnect"
	"main/src/httpHandler"
	"net/http"
	"os"
	"time"
)

func main() {
	loadEnv()
	time.Sleep(time.Second * 5)
	dbConnect.Init()

	http.HandleFunc("/add", httpHandler.AddNewTodo)
	http.HandleFunc("/delete", httpHandler.DeleteTodo)
	http.HandleFunc("/check", httpHandler.CheckTodo)
	http.HandleFunc("/get-all", httpHandler.GetAllTodos)

	err := http.ListenAndServeTLS(":"+os.Getenv("APP_PORT"), os.Getenv("APP_CERT_PATH"), os.Getenv("APP_KEY_PATH"), nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
}
