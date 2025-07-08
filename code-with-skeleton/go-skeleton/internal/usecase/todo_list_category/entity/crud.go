package entity

type TodoListCategoryReq struct {
	ID          int64  `json:"id,omitempty" swaggerignore:"true"`
	Name        string `json:"name" validate:"required" name:"Nama Kategori"`
	Description string `json:"description,omitempty" name:"Deskripsi"`
}

type TodoListCategoryResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

func (r *TodoListCategoryReq) SetID(id int64) {
	r.ID = id
}

func (r *TodoListCategoryReq) SetName(name string) {
	r.Name = name
}
