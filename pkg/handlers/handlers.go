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
