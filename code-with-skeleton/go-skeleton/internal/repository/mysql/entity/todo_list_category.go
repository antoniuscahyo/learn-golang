package entity

import "time"

type TodoListCategory struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedBy   int64     `gorm:"column:created_by" json:"created_by"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}

type TodoListCategoryResponse struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}

func (TodoListCategory) TableName() string {
	return "todo_list_categories"
}
