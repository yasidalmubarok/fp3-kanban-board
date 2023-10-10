package dto

import "time"

// ======= POST ========
type NewCategoryRequest struct {
	Type string `json:"type" valid:"required~type cannot be empty"`
}

type NewCategoryResponse struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}