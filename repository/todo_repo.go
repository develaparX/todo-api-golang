package repository

import (
	"database/sql"
	"math"
	"todo-api/models"
	"todo-api/models/dto"
)

type TodoRepository interface {
	CreateTodo(todo *models.Todo) error
	GetTodoByID(id string) (*models.Todo, error)
	GetAllTodos(page int, size int) ([]*models.Todo, dto.Paging, error)
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id string) error
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(database *sql.DB) TodoRepository {
	return &todoRepository{
		db: database,
	}
}

func (r *todoRepository) CreateTodo(todo *models.Todo) error {
	query := `INSERT INTO trx_todos (title, content, user_id) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	todoCreate := r.db.QueryRow(query, todo.Title, todo.Content, todo.User.ID).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)
	return todoCreate
}

func (r *todoRepository) GetTodoByID(id string) (*models.Todo, error) {
	todo := &models.Todo{}
	query := `SELECT t.id, t.title, t.content, t.created_at, t.updated_at, u.id, u.fullname, u.email, u.created_at, u.updated_at, u.role FROM trx_todos t JOIN mst_users u ON t.user_id = u.id WHERE t.id = $1`
	err := r.db.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Content, &todo.CreatedAt, &todo.UpdatedAt, &todo.User.ID, &todo.User.Fullname, &todo.User.Email, &todo.User.CreatedAt, &todo.User.UpdatedAt, &todo.User.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return todo, nil
}

func (r *todoRepository) GetAllTodos(page int, size int) ([]*models.Todo, dto.Paging, error) {

	skip := (page - 1) * size

	query := `SELECT t.id, t.title, t.content, t.created_at, t.updated_at, u.id, u.fullname, u.email, u.created_at, u.updated_at, u.role FROM trx_tasks t JOIN mst_users u ON t.user_id = u.id LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, size, skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	totalRows := 0
	err = r.db.QueryRow(`SELECT COUNT(*) FROM trx_tasks`).Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	defer rows.Close()

	var todos []*models.Todo
	for rows.Next() {
		todo := &models.Todo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.CreatedAt, &todo.UpdatedAt, &todo.User.ID, &todo.User.Fullname, &todo.User.Email, &todo.User.CreatedAt, &todo.User.UpdatedAt, &todo.User.Role)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		todos = append(todos, todo)
	}
	paging := dto.Paging{
		Page:       page,
		Size:       size,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64((totalRows) / (size)))),
	}
	return todos, paging, nil
}

func (r *todoRepository) UpdateTodo(todo *models.Todo) error {
	query := `UPDATE trx_todos SET title = $1, content = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3`
	_, err := r.db.Exec(query, todo.Title, todo.Content, todo.ID)
	return err
}

func (r *todoRepository) DeleteTodo(id string) error {
	query := `DELETE FROM trx_todos WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
