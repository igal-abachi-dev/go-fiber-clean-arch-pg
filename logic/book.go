package logic

import (
	"go-fiber-clean-arch-pg/entities"
	"go-fiber-clean-arch-pg/repositories"
)

type BookLogic interface {
	GetBookById(id int32) (entities.Book, error)
	GetBooksByCategory(category int32) ([]entities.Book, error)
}

type bookLogic struct {
	repository repositories.BookRepository
}

func NewBookLogic(r repositories.BookRepository) BookLogic {
	return &bookLogic{
		repository: r,
	}
}
func (l *bookLogic) GetBooksByCategory(category int32) ([]entities.Book, error) {
	items, _ := l.repository.GetBooksByCategory(category)
	return items, nil
}

func (l *bookLogic) GetBookById(id int32) (entities.Book, error) {
	item, _ := l.repository.GetBookById(id)
	return item, nil
}
