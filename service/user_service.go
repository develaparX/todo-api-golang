package service

import (
	"todo-api/models"
	"todo-api/models/dto"
	"todo-api/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetAllUsers(page int, size int) ([]*models.User, dto.Paging, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.userRepository.GetUserByID(id)
}

func (s *userService) GetAllUsers(page int, size int) ([]*models.User, dto.Paging, error) {
	return s.userRepository.GetAllUsers(page, size)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepository.UpdateUser(user)
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepository.DeleteUser(id)
}
