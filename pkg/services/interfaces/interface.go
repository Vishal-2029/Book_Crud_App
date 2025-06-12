package interfaces

import "github.com/Vishal-2029/BookCrud-App/pkg/model"

type BookserviceInter interface {
	Create(book *model.Books) error
	GetAll() ([]model.Books, error)
	Update(book *model.Books) error
	Delete(id int) error
}