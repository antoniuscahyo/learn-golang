package handler

import (
	"strconv"
	"todo-clean-arch/entity"
	"todo-clean-arch/usecase"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	Usecase *usecase.CategoryUsecase
}

func NewCategoryHandler(app *fiber.App, uc *usecase.CategoryUsecase) {
	handler := &CategoryHandler{Usecase: uc}
	app.Get("/categories", handler.GetAll)
	app.Get("/categories/:id", handler.GetByID)
	app.Post("/categories", handler.Create)
	app.Put("/categories/:id", handler.Update)
	app.Delete("/categories/:id", handler.Delete)
}

func (h *CategoryHandler) GetAll(c *fiber.Ctx) error {
	categories, err := h.Usecase.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(categories)
}

func (h *CategoryHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	category, err := h.Usecase.GetByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	return c.JSON(category)
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	var input entity.Category
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	category, err := h.Usecase.Create(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(category)
}

func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var input entity.Category
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	input.ID = uint(id)
	if err := h.Usecase.Update(input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(input)
}

func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.Usecase.Delete(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
