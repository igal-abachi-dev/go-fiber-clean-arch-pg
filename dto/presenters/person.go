package presenters

import (
	"go-fiber-clean-arch-pg/entities"
	"time"
)

type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func PersonFromEntity(e entities.Person) Person {
	return Person{
		ID:        e.ID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
	}
}

func PersonToEntity(p Person) entities.Person {
	return entities.Person{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
