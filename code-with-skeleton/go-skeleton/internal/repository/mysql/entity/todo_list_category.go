package entity

import "time"

type TodoListCategory struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}

func (TodoListCategory) TableName() string {
	return "todo_list_categories"
}
