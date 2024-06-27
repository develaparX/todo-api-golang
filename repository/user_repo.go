package repository

import (
	"database/sql"
	"math"
	"todo-api/models"
	"todo-api/models/dto"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetAllUsers(page int, size int) ([]*models.User, dto.Paging, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO mst_users (fullname, email, passwords, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	userCreated := r.db.QueryRow(query, user.Fullname, user.Email, user.Password, user.Role).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return userCreated
}

func (r *userRepository) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, fullname, email, created_at, updated_at, role FROM mst_users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Fullname, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetAllUsers(page int, size int) ([]*models.User, dto.Paging, error) {

	skip := (page - 1) * size

	query := `SELECT id, fullname, email, created_at, updated_at, role FROM mst_users LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, size, skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	totalRows := 0
	err = r.db.QueryRow(`SELECT COUNT(*) FROM mst_users`).Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Fullname, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Role)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		users = append(users, user)
	}
	paging := dto.Paging{
		Page:       page,
		Size:       size,
		TotalRows:  totalRows,
		TotalPages: int(math.Ceil(float64((totalRows) / (size)))),
	}
	return users, paging, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	query := `UPDATE mst_users SET fullname = $1, email = $2, role = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4`
	_, err := r.db.Exec(query, user.Fullname, user.Email, user.Role, user.ID)
	return err
}

func (r *userRepository) DeleteUser(id string) error {
	query := `DELETE FROM mst_users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
