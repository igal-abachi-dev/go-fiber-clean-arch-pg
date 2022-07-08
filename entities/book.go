package entities

type Book struct {
	Name       string  `json:"name" db:"Name"`
	Author     *string `json:"author" db:"Author"`
	Id         int32   `json:"id" db:"Id"` //`db:"book_id"`
	Category   int32   `json:"category" db:"Category"`
	Commentary *string `json:"commentary" db:"commentary"`
}
