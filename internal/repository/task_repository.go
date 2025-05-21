package repository

import (
	"todo-app/internal/domain"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *domain.Task) error
	GetByUserID(userID uint) ([]domain.Task, error)
	GetByID(id uint) (*domain.Task, error)
	Update(task *domain.Task) error
	Delete(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *domain.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetByUserID(userID uint) ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetByID(id uint) (*domain.Task, error) {
	var task domain.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *taskRepository) Update(task *domain.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Task{}, id).Error
}