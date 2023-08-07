package dbConnect

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	dsn = "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("SSL_MODE")
)

var TodoCounter uint
var db *gorm.DB

func Init() {
	loadEnv()
	var err error
	fmt.Println(dsn)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	TodoCounter = getGreaterId()
	TodoCounter++
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
}

func getGreaterId() uint {
	var greater int

	err := db.Model(&Todo{}).Select("MAX(id) as max_id").Scan(&greater).Error
	if err != nil {
		log.Fatal(err)
	}

	return uint(greater)
}

func GetAllTodos() []byte {
	var todos []Todo

	err := db.Find(&todos).Error
	if err != nil {
		log.Fatal("Failed to retrieve users: ", err)
	}

	todosJSON, err := json.Marshal(todos)
	if err != nil {
		log.Fatal("Failed to convert todos to JSON: ", err)
	}

	fmt.Println(string(todosJSON))

	return todosJSON
}

func AddTodo(title string) {
	todo := Todo{ID: TodoCounter, Title: title, Status: false}

	err := db.Create(&todo).Error
	if err != nil {
		log.Fatal(err)
	}

	TodoCounter++
}

func DeleteTodo(id int) {
	var todo Todo

	db.First(&todo, id)
	db.Delete(&todo)
}

func CheckTodo(id int) {
	var todo Todo
	db.First(&todo, id)

	todo.Status = !todo.Status
	db.Save(&todo)
}
