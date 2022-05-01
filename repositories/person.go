package repositories

import (
	"go-fiber-clean-arch-pg/entities"
)

//Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	GetAll() (*[]entities.Person, error)
	GetItem(ID string) (*entities.Person, error)
	//Create(p *entities.Person) (*entities.Person, error)
	//Update(p *entities.Person) (*entities.Person, error)
	//Delete(ID string) error
}
type repository struct {
	Collection map[string]*entities.Person
}

//NewRepo is the single instance repo that is being created.
func NewRepo() Repository {
	celebs := map[string]*entities.Person{
		"Nicolas Cage":       {ID: "1", FirstName: "Nicolas", LastName: "Cage"},
		"Selena Gomez":       {ID: "2", FirstName: "Selena", LastName: "Gomez"},
		"Scarlett Johansson": {ID: "3", FirstName: "Scarlett", LastName: "Johansson"},
	}
	var repo = &repository{celebs}
	return repo
}

//ReadBook is a mongo repository that helps to fetch books
func (r *repository) GetAll() (*[]entities.Person, error) {
	var books []entities.Person

	for _, value := range r.Collection {
		books = append(books, *value)
	}

	//cursor, err := r.Collection.Find(context.Background(), bson.D{})
	//if err != nil {
	//	return nil, err
	//}
	//for cursor.Next(context.TODO()) {
	//	var book entities.Person
	//	_ = cursor.Decode(&book)
	//	books = append(books, book)
	//}

	return &books, nil
}

func (r *repository) GetItem(ID string) (*entities.Person, error) {
	items := r.Collection
	for _, item := range items {
		if item.ID == ID {
			return item, nil
		}
	}
	return nil, nil
}

/*
//CreateBook is a mongo repository that helps to create books
func (r *repository) Create(book *entities.Person) (*entities.Person, error) {
	book.ID = primitive.NewObjectID()
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

//UpdateBook is a mongo repository that helps to update books
func (r *repository) Update(book *entities.Person) (*entities.Person, error) {
	book.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": book.ID}, bson.M{"$set": book})
	if err != nil {
		return nil, err
	}
	return book, nil
}

//DeleteBook is a mongo repository that helps to delete books
func (r *repository) Delete(ID string) error {
	bookID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": bookID})
	if err != nil {
		return err
	}
	return nil
}*/
