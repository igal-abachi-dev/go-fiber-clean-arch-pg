package repositories

import (
	"go-fiber-clean-arch-pg/entities"
)

//pgxpool: https://github.com/jackc/pgx
//pgxscan: https://github.com/georgysavva/scany
//https://play.sqlc.dev/
//

//Category:getAll()
//Book:getbycategory[],getitem(id)
//filecontent: getlinkbybook,getTextbyBook

/*
SELECT Name, Level, id FROM public."Category"
ORDER BY id ASC

SELECT Name, Author, Id, Category FROM public."Book"
WHERE "Category" = $1
ORDER BY "Id" ASC

SELECT Name, Author, Id, Category FROM public."Book"
WHERE "Id" = $1 LIMIT 1

SELECT url, password FROM public."FileContent"
WHERE "book_id" = $1 LIMIT 1;

SELECT content FROM public."FileContent"
WHERE "book_id" = $1 LIMIT 1;
*/
//search text in all books?

//or
//getall,getitem,getbyparent

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
func NewPersonRepo() Repository {
	celebs := map[string]*entities.Person{
		"": {ID: "1", FirstName: "", LastName: ""},
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
