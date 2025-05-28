package services

import (
    "errors"
    "todo-app/internal/domain"
    "todo-app/internal/repository"
)

type CategoryService interface {
    CreateCategory(category *domain.Category) error
    GetUserCategories(userID uint) ([]domain.Category, error)
    DeleteCategory(userID, categoryID uint) error
}

type categoryService struct {
    categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
    return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) CreateCategory(category *domain.Category) error {
    return s.categoryRepo.Create(category)
}

func (s *categoryService) GetUserCategories(userID uint) ([]domain.Category, error) {
    return s.categoryRepo.GetByUserID(userID)
}

func (s *categoryService) DeleteCategory(userID, categoryID uint) error {
    category, err := s.categoryRepo.GetByID(categoryID)
    if err != nil {
        return err
    }

    if category.UserID != userID {
        return errors.New("category does not belong to user")
    }

    return s.categoryRepo.Delete(categoryID)
}