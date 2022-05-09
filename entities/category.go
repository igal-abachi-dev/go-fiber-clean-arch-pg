package entities

type Category struct {
	Name  string `json:"name"`
	Level *int32 `json:"level"`
	ID    int32  `json:"id"`
}

//type NullInt32 struct {
//	Int32 int32 //*int32
//	Valid bool // Valid is true if Int32 is not NULL
//}
