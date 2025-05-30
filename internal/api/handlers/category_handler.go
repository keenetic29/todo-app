package handlers

import (
	"net/http"
	"strconv"
	"todo-app/internal/domain"
	"todo-app/internal/services"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
    categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) *CategoryHandler {
    return &CategoryHandler{categoryService: categoryService}
}

// CreateCategory godoc
// @Summary Создать новую категорию
// @Description Создает новую категорию для авторизованного пользователя
// @Tags Категории
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body domain.Category true "Данные категории"
// @Success 201 {object} domain.SwaggerCategory
// @Failure 400 {object} domain.ErrorResponse
// @Failure 401 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
    userID := c.MustGet("userID").(uint)

    var category domain.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    category.UserID = userID
    if err := h.categoryService.CreateCategory(&category); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

     // Преобразуем для свагера
    swaggerCategory := domain.SwaggerCategory{
        ID:        category.ID,
        Name:      category.Name,
        UserID:    category.UserID,
        CreatedAt: category.CreatedAt,
        UpdatedAt: category.UpdatedAt,
    }
    
    c.JSON(http.StatusCreated, swaggerCategory)
}

// GetCategories godoc
// @Summary Получить категории пользователя
// @Description Возвращает все категории для авторизованного пользователя
// @Tags Категории
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} domain.SwaggerCategory
// @Failure 401 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
    userID := c.MustGet("userID").(uint)

    categories, err := h.categoryService.GetUserCategories(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Преобразуем для свагера
    var result []domain.SwaggerCategory
    for _, category := range categories {
        result = append(result, domain.SwaggerCategory{
            ID:        category.ID,
            Name:      category.Name,
            UserID:    category.UserID,
            CreatedAt: category.CreatedAt,
            UpdatedAt: category.UpdatedAt,
        })
    }

    c.JSON(http.StatusOK, result)
}

// DeleteCategory godoc
// @Summary Удалить категорию
// @Description Удаляет категорию по ID (должна принадлежать пользователю)
// @Tags Категории
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID категории"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 401 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
    userID := c.MustGet("userID").(uint)
    categoryID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
        return
    }

    if err := h.categoryService.DeleteCategory(userID, uint(categoryID)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Category deleted successfully"})
}