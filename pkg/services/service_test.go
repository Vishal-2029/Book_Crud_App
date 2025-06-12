package services_test

import (
	"testing"

	"github.com/Vishal-2029/BookCrud-App/pkg/model"
	"github.com/Vishal-2029/BookCrud-App/pkg/services"
	"github.com/stretchr/testify/assert"
)

type TestBookservice struct{}

func (t *TestBookservice) Create(book *model.Books) error {
	return nil
}

func (t *TestBookservice) GetAll() ([]model.Books, error) {
	return []model.Books{
		{
			ID:            1,
			Title:         "Test Book",
			Author:        "Test Author",
			PublishedYear: 2023,
		},
	}, nil
}

func (t *TestBookservice) Update(book *model.Books) error {
	return nil
}

func (t *TestBookservice) Delete(id int) error {
	return nil
}

// TestBookservice_Create tests the creation of a new book using the BookSVC service.
func TestBookservice_Create(t *testing.T) {
	testBook := &TestBookservice{}
	BookSVC := services.NewBookservices(testBook)
	NewBook := &model.Books{
		Title:         "Test Book",
		Author:        "Test Author",
		PublishedYear: 2024,
	}
	err := BookSVC.Create(NewBook)
	assert.NoError(t, err, "Expected no error when creating a book")
}

func TestBookservice_GetAll(t *testing.T) {
	testBook := &TestBookservice{}
	BookSVC := services.NewBookservices(testBook)

	books, err := BookSVC.GetAll()
	assert.NoError(t, err, "Expected no error when getting all books")
	assert.NotEmpty(t, books, "Expected books to not be empty")
	assert.Equal(t, 1, len(books), "Expected one book in the list")
	assert.Equal(t, "Test Book", books[0].Title, "Expected book title to match")
}


func TestBookservice_Update(t *testing.T) {
	testBook := &TestBookservice{}
	BookSVC := services.NewBookservices(testBook)

	updatedBook := &model.Books{
		ID: 		  1,
		Title:       "Updated Test Book",
		Author: 	"Updated Test Author",
		PublishedYear: 2025,
	}

	err := BookSVC.Update(updatedBook)
	assert.NoError(t, err, "Expected no error when updating a book")
}

func TestBookservice_Delete(t *testing.T) {
	testBook := &TestBookservice{}
	BookSVC := services.NewBookservices(testBook)

	err := BookSVC.Delete(1)
	assert.NoError(t, err, "Expected no error when deleting a book")
}