package services

import (
	"github.com/Vishal-2029/BookCrud-App/pkg/model"
	interfaces "github.com/Vishal-2029/BookCrud-App/pkg/repo/interface"
)

type BookService struct{
	repo interfaces.BookRepoInter
}

func (B *BookService) Create(book *model.Books) error{
	return B.repo.Create(book)
}	

func (B *BookService) GetAll() ([]model.Books,error){
	return B.repo.GetAll()
}

func (B *BookService)Update(book *model.Books) error {
	return B.repo.Update(book)
}

func (B *BookService) Delete(id int) error {
	return B. repo.Delete(id)
}

func NewBookservices(repo interfaces.BookRepoInter) *BookService {
	return &BookService{
		repo: repo,
	}
}
