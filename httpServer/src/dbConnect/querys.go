package dbConnect

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const dsn = "host=10.11.19.156 user=admin-poc password=poc-ssip dbname=todolist port=30432 sslmode=disable"

func GetAllTodos() []byte {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var todos []Todo

	err = db.Find(&todos).Error
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	todo := Todo{ID: 1, Title: title, Status: false}

	err = db.Create(&todo).Error
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteTodo(id int) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var todo Todo
	db.First(&todo, id)

	db.Delete(&todo)
}

func CheckTodo(id int) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var todo Todo
	db.First(&todo, id)

	todo.Status = !todo.Status
	db.Save(&todo)
}
