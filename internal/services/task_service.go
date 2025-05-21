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
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{taskRepo: taskRepo}
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