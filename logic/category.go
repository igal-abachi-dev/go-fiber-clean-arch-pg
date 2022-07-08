package logic

import (
	"go-fiber-clean-arch-pg/entities"
	"go-fiber-clean-arch-pg/repositories"
)

type CategoryLogic interface {
	GetAllCategories() ([]entities.Category, error)
}

type categoryLogic struct {
	repository repositories.CategoryRepository
}

func NewCategoryLogic(r repositories.CategoryRepository) CategoryLogic {
	return &categoryLogic{
		repository: r,
	}
}
func (l *categoryLogic) GetAllCategories() ([]entities.Category, error) {
	items, _ := l.repository.GetAllCategories()

	return items, nil
}
