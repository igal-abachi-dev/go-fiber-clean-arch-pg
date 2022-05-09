package entities

type Book struct {
	Name     string  `json:"name"`
	Author   *string `json:"author"`
	Id       int32   `json:"id"` //`db:"book_id"`
	Category int32   `json:"category"`
}
