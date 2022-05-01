package logic

import (
	"go-fiber-clean-arch-pg/dto/presenters"
	"go-fiber-clean-arch-pg/repositories"
)

//Service is an interface from which our api module can access our repository of all our models
type PersonLogic interface {
	GetAll() (*[]presenters.Person, error)
	GetItem(ID string) (*presenters.Person, error)
	//InsertPerson(book *presenter.Person) (*string, error)
	//UpdatePerson(book *presenter.Person) error
	//RemovePerson(ID string) error
}

type logic struct {
	repository repositories.Repository
}

//NewService is used to create a single instance of the service
func NewService(r repositories.Repository) PersonLogic {
	return &logic{
		repository: r,
	}
}

//https://github.com/go-playground/validator

//https://github.com/francoispqt/gojay
//json.NewEncoder(httpResponseWriter).Encode(&obj)

//https://github.com/jackc/pgx
// https://github.com/georgysavva/scany

//pgxpool  pgxscan
//https://play.sqlc.dev/

// https://github.com/uber-go/zap

//https://github.com/sarulabs/di#scopes-in-practice
//https://github.com/google/wire
//maybe don't use di is better

// https://github.com/jessevdk/go-flags

func (s *logic) GetAll() (*[]presenters.Person, error) { //flat items
	itemsRef, _ := s.repository.GetAll()
	items := *itemsRef
	persons := []presenters.Person{}
	for i := 0; i < len(items); i++ {
		item := items[i]
		persons = append(persons, presenters.PersonFromEntity(item))
	}
	return &persons, nil
}

//https://winterflower.github.io/2017/08/20/the-asterisk-and-the-ampersand/

func (s *logic) GetItem(ID string) (*presenters.Person, error) { //full item
	item, _ := s.repository.GetItem(ID)
	person := presenters.PersonFromEntity(*item)
	return &person, nil
}

//
////InsertBook is a service layer that helps insert book in BookShop
//func (s *logic) InsertPerson(book *entities.Book) (*entities.Book, error) {
//	return s.repository.CreateBook(book)
//}
//
//
////UpdateBook is a service layer that helps update books in BookShop
//func (s *logic) UpdatePerson(book *entities.Book) (*entities.Book, error) {
//	return s.repository.UpdatePerson(book)
//}
//
////RemovePerson is a service layer that helps remove books from BookShop
//func (s *logic) RemovePerson(ID string) error {
//	return s.repository.DeletePerson(ID)
//}
