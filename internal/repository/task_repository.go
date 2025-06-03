package repository

import (
	"fmt"
	"strings"
	"todo-app/internal/domain"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *domain.Task) error
	GetByUserID(userID uint, query domain.TaskQuery) ([]domain.Task, int64, error)
	GetByID(id uint) (*domain.Task, error)
	Update(task *domain.Task) error
	Delete(id uint) error
	UpdateCategory(taskID uint, categoryID *uint) error
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

func (r *taskRepository) GetByUserID(userID uint, query domain.TaskQuery) ([]domain.Task, int64, error) {
    var tasks []domain.Task
    var total int64

    db := r.db.Model(&domain.Task{}).Where("user_id = ?", userID)

    // Фильтрация по статусу
    if query.Completed != nil {
        db = db.Where("completed = ?", *query.Completed)
    }

    // Сортировка
    if query.SortBy != "" {
        order := "ASC"
        sortField := query.SortBy

        if strings.HasPrefix(query.SortBy, "-") {
            order = "DESC"
            sortField = strings.TrimPrefix(query.SortBy, "-")
        }

        // Проверяем допустимые поля для сортировки
        validSortFields := map[string]bool{
            "created_at":  true,
            "updated_at":  true,
            "title":       true,
            "completed":  true,
        }

        if validSortFields[sortField] {
            db = db.Order(fmt.Sprintf("%s %s", sortField, order))
        }
    }

    // Получаем общее количество
    err := db.Count(&total).Error
    if err != nil {
        return nil, 0, err
    }

    // Применяем пагинацию
    if query.Page < 1 {
        query.Page = 1
    }
    if query.Limit < 1 || query.Limit > 100 {
        query.Limit = 10
    }

    offset := (query.Page - 1) * query.Limit
    err = db.Offset(offset).Limit(query.Limit).Find(&tasks).Error

    return tasks, total, err
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

func (r *taskRepository) UpdateCategory(taskID uint, categoryID *uint) error {
    return r.db.Model(&domain.Task{}).Where("id = ?", taskID).Update("category_id", categoryID).Error
}