package services

import (
	"errors"
	"todo-app/internal/domain"
	"todo-app/internal/repository"
)

type TaskService interface {
	CreateTask(task *domain.Task) error
	GetUserTasks(userID uint) ([]domain.Task, error)
	GetTaskByID(userID, taskID uint) (*domain.Task, error)
	UpdateTask(userID uint, task *domain.Task) error
	DeleteTask(userID, taskID uint) error
	UpdateTaskCategory(userID, taskID uint, categoryID *uint) error
}

type taskService struct {
	taskRepo repository.TaskRepository
	categoryRepo repository.CategoryRepository
}

func NewTaskService(taskRepo repository.TaskRepository, categoryRepo repository.CategoryRepository) TaskService {
	return &taskService{
		taskRepo:     taskRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *taskService) CreateTask(task *domain.Task) error {
	return s.taskRepo.Create(task)
}

func (s *taskService) GetUserTasks(userID uint) ([]domain.Task, error) {
	return s.taskRepo.GetByUserID(userID)
}

func (s *taskService) GetTaskByID(userID, taskID uint) (*domain.Task, error) {
	task, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		return nil, err
	}

	if task.UserID != userID {
		return nil, errors.New("task does not belong to user")
	}

	return task, nil
}

func (s *taskService) UpdateTask(userID uint, task *domain.Task) error {
	existingTask, err := s.GetTaskByID(userID, task.ID)
	if err != nil {
		return err
	}

	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.Completed = task.Completed

	return s.taskRepo.Update(existingTask)
}

func (s *taskService) DeleteTask(userID, taskID uint) error {
	_, err := s.GetTaskByID(userID, taskID)
	if err != nil {
		return err
	}

	return s.taskRepo.Delete(taskID)
}

func (s *taskService) UpdateTaskCategory(userID, taskID uint, categoryID *uint) error {
	// Проверяем что задача принадлежит пользователю
	task, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		return domain.ErrTaskNotFound
	}
	
	if task.UserID != userID {
		return domain.ErrUnauthorized
	}

	// Если передана категория, проверяем что она существует и принадлежит пользователю
	if categoryID != nil {
		category, err := s.categoryRepo.GetByID(*categoryID)
		if err != nil {
			return domain.ErrCategoryNotFound
		}
		if category.UserID != userID {
			return domain.ErrUnauthorized
		}
	}

	// Обновляем категорию задачи
	return s.taskRepo.UpdateCategory(taskID, categoryID)
}