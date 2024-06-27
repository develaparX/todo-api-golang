package controller

import (
	"net/http"
	"strconv"
	"todo-api/models"
	"todo-api/models/dto"
	"todo-api/service"
	"todo-api/utils"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService service.TodoService
	rg          *gin.RouterGroup
}

func (t *TodoController) Route() {
	router := t.rg.Group("/todo")
	{
		router.GET("/", t.GetAllTodos)
		router.POST("/", t.CreateTodo)
		router.PUT("/:id", t.UpdateTodo)
		router.DELETE("/:id", t.DeleteTodo)
		router.GET("/:id", t.GetTodoByID)
	}
}

func NewTodoController(todoService service.TodoService, rg *gin.RouterGroup) *TodoController {
	return &TodoController{
		todoService: todoService,
		rg:          rg,
	}
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var req dto.CreateTodoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	todo := &models.Todo{
		Title:   req.Title,
		Content: req.Content,
		User: models.User{
			ID: req.UserID,
		},
	}

	if err := c.todoService.CreateTodo(todo); err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendSingleResponse(ctx, "Todo created successfully", todo, http.StatusCreated)
}

func (c *TodoController) GetTodoByID(ctx *gin.Context) {
	todoID := ctx.Param("id")

	todo, err := c.todoService.GetTodoByID(todoID)
	if err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if todo == nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusNotFound)
		return
	}

	utils.SendSingleResponse(ctx, "Todo retrieved successfully", todo, http.StatusOK)
}

func (c *TodoController) GetAllTodos(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
	}

	size, err2 := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err2 != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
	}

	data, paging, err := c.todoService.GetAllTodos(page, size)
	if err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
	}

	var listData []interface{}

	for _, pl := range data {
		listData = append(listData, pl)
	}

	utils.SendPagingResponse(ctx, "Success Get Data", listData, paging, http.StatusOK)
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	todoID := ctx.Param("id")

	var req dto.UpdateTodoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := c.todoService.GetTodoByID(todoID)
	if err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if todo == nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusNotFound)
		return
	}

	todo.Title = req.Title
	todo.Content = req.Content

	if err := c.todoService.UpdateTodo(todo); err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendSingleResponse(ctx, "Todo updated successfully", todo, http.StatusOK)
}

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	todoID := ctx.Param("id")

	err := c.todoService.DeleteTodo(todoID)
	if err != nil {
		utils.SendErrorResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendSingleResponse(ctx, "Todo deleted successfully", nil, http.StatusOK)
}
