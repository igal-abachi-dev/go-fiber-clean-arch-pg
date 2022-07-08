package entities

type Category struct {
	Name  string `json:"name" db:"Name"`
	Level *int32 `json:"level" db:"Level"`
	//	id    *int32 `json:"id"  db:"id"`
}

//type NullInt32 struct {
//	Int32 int32 //*int32
//	Valid bool // Valid is true if Int32 is not NULL
//}
