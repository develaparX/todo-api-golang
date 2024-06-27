package service

import (
	"todo-api/models"
	"todo-api/models/dto"
	"todo-api/repository"
)

type TodoService interface {
	CreateTodo(todo *models.Todo) error
	GetTodoByID(id string) (*models.Todo, error)
	GetAllTodos(page int, size int) ([]*models.Todo, dto.Paging, error)
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id string) error
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepo,
	}
}

func (s *todoService) CreateTodo(todo *models.Todo) error {
	return s.todoRepository.CreateTodo(todo)
}

func (s *todoService) GetTodoByID(id string) (*models.Todo, error) {
	return s.todoRepository.GetTodoByID(id)
}

func (s *todoService) GetAllTodos(page int, size int) ([]*models.Todo, dto.Paging, error) {
	return s.todoRepository.GetAllTodos(page, size)
}

func (s *todoService) UpdateTodo(todo *models.Todo) error {
	return s.todoRepository.UpdateTodo(todo)
}

func (s *todoService) DeleteTodo(id string) error {
	return s.todoRepository.DeleteTodo(id)
}
