package dto

import "time"

// ======= POST ========
type NewCategoryRequest struct {
	Type string `json:"type" valid:"required~type cannot be empty"`
}

type NewCategoryResponse struct {
	Id        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// ====== GET ========
type GetCategoryResponse struct {
	Id        uint        `json:"id"`
	Type      string      `json:"type"`
	UpdatedAt time.Time   `json:"updated_at"`
	CreatedAt time.Time   `json:"created_at"`
	Tasks     []TaskDatas `json:"Tasks"`
}
