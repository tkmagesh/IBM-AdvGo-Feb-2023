package models

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

/* GORM expectation */
func (t *Todo) TableName() string {
	return "todo"
}
