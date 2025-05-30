package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"todo-app/internal/domain"
	"todo-app/internal/services"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

// GetTasks godoc
// @Summary Получить все задачи пользователя
// @Description Возвращает список всех задач для авторизованного пользователя
// @Tags Задачи
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} domain.SwaggerTask
// @Failure 401 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /tasks [get]
func (h *TaskHandler) GetTasks(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	tasks, err := h.taskService.GetUserTasks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// для свагера
	var result []domain.SwaggerTask
	for _, task := range tasks {
		result = append(result, task.ToSwagger())
	}

	//возвращаем не tasks, а result
	c.JSON(http.StatusOK, result)
}

// GetTaskByID godoc
// @Summary Получить задачу по ID
// @Description Возвращает задачу по указанному ID для авторизованного пользователя
// @Tags Задачи
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} domain.SwaggerTask
// @Failure 400 {object} domain.ErrorResponse
// @Failure 401 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
    userID := c.MustGet("userID").(uint)
    taskID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
        return
    }

    task, err := h.taskService.GetTaskByID(userID, uint(taskID))
    if err != nil {
        if errors.Is(err, domain.ErrTaskNotFound) || errors.Is(err, domain.ErrUnauthorized) {
            c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	// преобразуем для свагера
    c.JSON(http.StatusOK, task.ToSwagger())
}

// CreateTask godoc
// @Summary Создать новую задачу
// @Description Создает новую задачу для авторизованного пользователя
// @Tags Задачи
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body domain.Task true "Данные задачи"
// @Success 201 {object} domain.SwaggerTask
// @Failure 400 {object} domain.ErrorResponse
// @Failure 401 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.UserID = userID
	if err := h.taskService.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// преобразуем для свагера
	c.JSON(http.StatusCreated, task.ToSwagger())
}

// UpdateTask godoc
// @Summary Обновить задачу
// @Description Обновляет существующую задачу
// @Tags Задачи
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param input body domain.Task true "Обновленные данные задачи"
// @Success 200 {object} domain.SwaggerTask
// @Failure 400 {object} domain.ErrorResponse
// @Failure 401 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = uint(taskID)
	if err := h.taskService.UpdateTask(userID, &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// преобразуем для свагера
	c.JSON(http.StatusCreated, task.ToSwagger())
}

// DeleteTask godoc
// @Summary Удалить задачу
// @Description Удаляет задачу по указанному ID
// @Tags Задачи
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 401 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	if err := h.taskService.DeleteTask(userID, uint(taskID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Task deleted successfully"})
}

// UpdateTaskCategory godoc
// @Summary Обновить категорию задачи
// @Description Обновляет или удаляет категорию для задачи
// @Tags Задачи
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param input body domain.CategoryRequest true "Данные категории"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 401 {object} domain.ErrorResponse
// @Failure 403 {object} domain.ErrorResponse
// @Failure 404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /tasks/{id}/category [patch]
func (h *TaskHandler) UpdateTaskCategory(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var request domain.CategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.taskService.UpdateTaskCategory(userID, uint(taskID), request.CategoryID); err != nil {
		switch err {
		case domain.ErrTaskNotFound, domain.ErrCategoryNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case domain.ErrUnauthorized:
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Task category updated successfully"})
}