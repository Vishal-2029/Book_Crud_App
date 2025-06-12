package repo

import (
	"github.com/Vishal-2029/BookCrud-App/pkg/model"
	"gorm.io/gorm"
)

type BookRepo struct {
	DB *gorm.DB
}

func (B *BookRepo) Create(book *model.Books) error {
	if err := B.DB.Create(&book).Error; err != nil {
		return err
	}

	return nil
}

func (B *BookRepo) GetAll() ([]model.Books, error) {
	var Books []model.Books
	if err := B.DB.Find(&Books).Error; err != nil {
		return nil, err
	}
	return Books, nil
}

func (B *BookRepo) Update(book *model.Books) error {
	if err := B.DB.Save(&book).Error; err != nil {
		return err
	}
	return nil
}

func (B *BookRepo) Delete(id int) error {
	if err := B.DB.Delete(&model.Books{}, id).Error; err != nil {
		return err
	}
	return nil
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{
		DB: db,
	}
}
