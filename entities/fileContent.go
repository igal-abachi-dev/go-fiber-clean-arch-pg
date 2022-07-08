package entities

//type NullString struct {
//	String string //*string
//	Valid  bool   // Valid is true if String is not NULL
//}

type GetBookTextContentRow struct {
	Content *string `json:"content" db:"content"`
}

type GetBookUrlRow struct {
	Url      string  `json:"url" db:"url"`
	Password *string `json:"password" db:"password"`
}

//
//type FileContent struct {
//	Url      string  `json:"url"`
//	Password *string `json:"password"`
//	Content  *string `json:"content"`
//	ID       int32   `json:"id"`
//	BookID   int32   `json:"book_id"`
//}
