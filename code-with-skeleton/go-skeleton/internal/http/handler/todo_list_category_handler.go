package handler

import (
	"net/http"

	"github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/internal/parser"
	"github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/internal/presenter/json"
	todoCategoryUsecase "github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/internal/usecase/todo_list_category"
	"github.com/antoniuscahyo/learn-golang/code-with-skeleton/go-skeleton/internal/usecase/todo_list_category/entity"

	"github.com/gofiber/fiber/v2"
)

type TodoListCategoryHandler struct {
	parser    parser.Parser
	presenter json.JsonPresenter
	usecase   todoCategoryUsecase.ICrudTodoListCategoryUsecase
}

func NewTodoListCategoryHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	usecase todoCategoryUsecase.ICrudTodoListCategoryUsecase,
) *TodoListCategoryHandler {
	return &TodoListCategoryHandler{parser, presenter, usecase}
}

func (h *TodoListCategoryHandler) Register(app fiber.Router) {
	app.Get("/todo-list-categories", h.GetAll)
	app.Get("/todo-list-categories/:id", h.GetByID)
	app.Post("/todo-list-categories", h.Create)
	app.Put("/todo-list-categories/:id", h.Update)
	app.Delete("/todo-list-categories/:id", h.Delete)
}

// @Summary Get All Todo List Categories
// @Tags Todo List Category
// @Success 200 {object} entity.GeneralResponse{data=[]entity.TodoListCategoryResponse}
// @Router /api/v1/todo-list-categories [get]
func (h *TodoListCategoryHandler) GetAll(c *fiber.Ctx) error {
	data, err := h.usecase.GetAll(c.Context())
	if err != nil {
		return h.presenter.BuildError(c, err)
	}
	return h.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary Get Todo List Category by ID
// @Tags Todo List Category
// @Param id path int true "Category ID"
// @Success 200 {object} entity.GeneralResponse{data=entity.TodoListCategoryResponse}
// @Router /api/v1/todo-list-categories/{id} [get]
func (h *TodoListCategoryHandler) GetByID(c *fiber.Ctx) error {
	id, err := h.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	data, err := h.usecase.GetByID(c.Context(), id)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}
	return h.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary Create a new Todo List Category
// @Tags Todo List Category
// @Param req body entity.TodoListCategoryReq true "Request Body"
// @Success 201 {object} entity.GeneralResponse{data=entity.TodoListCategoryResponse}
// @Router /api/v1/todo-list-categories [post]
func (h *TodoListCategoryHandler) Create(c *fiber.Ctx) error {
	var req entity.TodoListCategoryReq

	if err := h.parser.ParserBodyRequest(c, &req); err != nil {
		return h.presenter.BuildError(c, err)
	}

	data, err := h.usecase.Create(c.Context(), req)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, data, "Created", http.StatusCreated)
}

// @Summary Update a Todo List Category
// @Tags Todo List Category
// @Param id path int true "Category ID"
// @Param req body entity.TodoListCategoryReq true "Request Body"
// @Success 200 {object} entity.GeneralResponse
// @Router /api/v1/todo-list-categories/{id} [put]
func (h *TodoListCategoryHandler) Update(c *fiber.Ctx) error {
	var req entity.TodoListCategoryReq

	if err := h.parser.ParserBodyWithIntIDPathParams(c, &req); err != nil {
		return h.presenter.BuildError(c, err)
	}

	if err := h.usecase.UpdateByID(c.Context(), req); err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, nil, "Updated", http.StatusOK)
}

// @Summary Delete a Todo List Category
// @Tags Todo List Category
// @Param id path int true "Category ID"
// @Success 200 {object} entity.GeneralResponse
// @Router /api/v1/todo-list-categories/{id} [delete]
func (h *TodoListCategoryHandler) Delete(c *fiber.Ctx) error {
	id, err := h.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	if err := h.usecase.DeleteByID(c.Context(), id); err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, nil, "Deleted", http.StatusOK)
}
