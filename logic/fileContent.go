package logic

import (
	"go-fiber-clean-arch-pg/entities"
	"go-fiber-clean-arch-pg/repositories"
)

type FileContentLogic interface {
	GetBookTextContent(bookID int32) (entities.GetBookTextContentRow, error)
	GetBookUrl(bookID int32) (entities.GetBookUrlRow, error)
}

type fileContentLogic struct {
	repository repositories.FileContentRepository
}

func NewFileContentLogic(r repositories.FileContentRepository) FileContentLogic {
	return &fileContentLogic{
		repository: r,
	}
}

func (l *fileContentLogic) GetBookTextContent(bookID int32) (entities.GetBookTextContentRow, error) {
	item, _ := l.repository.GetBookTextContent(bookID)
	return item, nil
}

func (l *fileContentLogic) GetBookUrl(bookID int32) (entities.GetBookUrlRow, error) {
	item, _ := l.repository.GetBookUrl(bookID)
	return item, nil
}
