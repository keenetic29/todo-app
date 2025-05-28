package repository

import (
    "todo-app/internal/domain"
    "gorm.io/gorm"
)

type CategoryRepository interface {
    Create(category *domain.Category) error
    GetByUserID(userID uint) ([]domain.Category, error)
    GetByID(id uint) (*domain.Category, error)
    Delete(id uint) error
}

type categoryRepository struct {
    db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
    return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *domain.Category) error {
    return r.db.Create(category).Error
}

func (r *categoryRepository) GetByUserID(userID uint) ([]domain.Category, error) {
    var categories []domain.Category
    err := r.db.Where("user_id = ?", userID).Find(&categories).Error
    return categories, err
}

func (r *categoryRepository) GetByID(id uint) (*domain.Category, error) {
    var category domain.Category
    err := r.db.First(&category, id).Error
    return &category, err
}

func (r *categoryRepository) Delete(id uint) error {
	// Обнуляем category_id у всех связанных задач
    if err := r.db.Model(&domain.Task{}).
        Where("category_id = ?", id).
        Update("category_id", nil).Error; err != nil {
        return err
    }


    return r.db.Delete(&domain.Category{}, id).Error
}