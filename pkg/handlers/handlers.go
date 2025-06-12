package handlers

import (
	"strconv"

	"github.com/Vishal-2029/BookCrud-App/pkg/model"
	"github.com/Vishal-2029/BookCrud-App/pkg/services/interfaces"
	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	SVC interfaces.BookserviceInter
}

func NewBookHandler(router fiber.Router,svc interfaces.BookserviceInter){
	handlers := &BookHandler{SVC: svc}

	router.Post("/books", handlers.Create)
	router.Get("/books", handlers.GetAll)
	router.Put("/books/:id", handlers.Update)
	router.Delete("/books/:id", handlers.Delete)
}


// @Summary Create a new book
// @Description Creates a new book with the provided data
// @Tags books
// @Accept json
// @Produce json
// @Param book body model.Books true "Book to create"
// @Success 201 {object} model.Books
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [post]
func (h *BookHandler) Create(c *fiber.Ctx) error {
	Book := new(model.Books)
	if err := c.BodyParser(Book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid request body",
		})
	}
	if err := h.SVC.Create(Book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{
			"error": "Failed to create book",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"message": "Book created successfully",
		"Data":    Book,
	})
}


// @Summary Get all books
// @Description Returns all available books
// @Tags books
// @Produce json
// @Success 200 {array} model.Books
// @Failure 500 {object} map[string]string
// @Router /books [get]
func (h *BookHandler) GetAll(c *fiber.Ctx) error {
	Books, err := h.SVC.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{
			"error": "Failed to retrieve books",
		})
	}
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message": "Books retrieved successfully",
		"Data":    Books,
	})
}

// @Summary Update a book
// @Description Updates the book with the given ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body model.Books true "Updated book"
// @Success 200 {object} model.Books
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [put]
func (h *BookHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": "ID is required",
		})
	}

	var updatedBook model.Books
	if err := c.BodyParser(&updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid request body",
		})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid ID format",
		})
	}
	updatedBook.ID = idInt // Ensure the ID is assigned

	if err := h.SVC.Update(&updatedBook); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{
			"error": "Failed to update book",
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message": "Book updated successfully",
		"Data":    updatedBook,
	})
}

// @Summary Delete a book
// @Description Deletes the book with the given ID
// @Tags books
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [delete]
func (h *BookHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": "ID is required",
		})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid ID format",
		})
	}

	if err := h.SVC.Delete(idInt); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{
			"error": "Failed to delete book",
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]string{
		"message": "Book deleted successfully",
	})
}
