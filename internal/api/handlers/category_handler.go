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

    c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
    userID := c.MustGet("userID").(uint)

    categories, err := h.categoryService.GetUserCategories(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, categories)
}

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

    c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}