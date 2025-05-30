package domain

import "time"

// SwaggerTask используется только для документации
type SwaggerTask struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Completed   bool             `json:"completed"`
	UserID      uint             `json:"user_id"`
	CategoryID  *uint            `json:"category_id,omitempty"`
	Category    *SwaggerCategory `json:"category,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

// SwaggerCategory используется только для документации
type SwaggerCategory struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Преобразование Task в SwaggerTask
func (t *Task) ToSwagger() SwaggerTask {
	res := SwaggerTask{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
		UserID:      t.UserID,
		CategoryID:  t.CategoryID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}

	if t.Category != nil {
		res.Category = &SwaggerCategory{
			ID:        t.Category.ID,
			Name:      t.Category.Name,
			UserID:    t.Category.UserID,
			CreatedAt: t.Category.CreatedAt,
			UpdatedAt: t.Category.UpdatedAt,
		}
	}

	return res
}


// ErrorResponse - стандартный формат ошибки
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse - стандартный формат успешного ответа
type SuccessResponse struct {
	Message string `json:"message"`
}