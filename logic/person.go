package logic

import "fiber-rest-api/entities"

//Service is an interface from which our api module can access our repository of all our models
type PersonLogic interface {
	InsertPerson(book *entities.Person) (*entities.Person, error)
	FetchPersons() (*[]presenter.Book, error)
	UpdatePerson(book *entities.Person) (*entities.Person, error)
	RemovePerson(ID string) error
}

type logic struct {
	repository Repository
}

//NewService is used to create a single instance of the service
func NewService(r Repository) PersonLogic {
	return &logic{
		repository: r,
	}
}

//InsertBook is a service layer that helps insert book in BookShop
func (s *logic) InsertPerson(book *entities.Book) (*entities.Book, error) {
	return s.repository.CreateBook(book)
}

//FetchBooks is a service layer that helps fetch all books in BookShop
func (s *logic) FetchPersons() (*[]presenter.Book, error) {
	return s.repository.ReadBook()
}

//UpdateBook is a service layer that helps update books in BookShop
func (s *logic) UpdatePerson(book *entities.Book) (*entities.Book, error) {
	return s.repository.UpdatePerson(book)
}

//RemovePerson is a service layer that helps remove books from BookShop
func (s *logic) RemovePerson(ID string) error {
	return s.repository.DeletePerson(ID)
}
