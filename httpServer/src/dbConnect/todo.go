package dbConnect

type Todo struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}