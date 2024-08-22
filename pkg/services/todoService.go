package services

import (
	"goWebService/pkg/models"
	"goWebService/pkg/utils"
)

type TodoService struct{}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) CreateTodo(todo *models.Todo) error {
	return utils.DB.Create(todo).Error
}

func (s *TodoService) GetAllTodos(todos *[]models.Todo) error {
	return utils.DB.Find(todos).Error
}

func (s *TodoService) GetTodoByID(todo *models.Todo, id string) error {
	return utils.DB.First(todo, id).Error
}

func (s *TodoService) GetTodosByUserId(todos *[]models.Todo, userId string) error {
	return utils.DB.Where("userId = ?", userId).Find(todos).Error
}

func (s *TodoService) UpdateTodo(todo *models.Todo) error {
	return utils.DB.Save(todo).Error
}

func (s *TodoService) DeleteTodoById(todo *models.Todo, id string) error {
	return utils.DB.Delete(id).Error
}

func (s *TodoService) DeleteUserTodos(userId string) error {
	return utils.DB.Where("userId = ?", userId).Delete(&models.Todo{}).Error
}
