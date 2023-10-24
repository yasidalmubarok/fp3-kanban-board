package dto

import "time"

// ======= POST ========
type NewCategoryRequest struct {
	Type string `json:"type" valid:"required~type cannot be empty"`
}

type NewCategoryResponse struct {
	Id        int       `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// ====== GET ========
type GetCategoryResponse struct {
	Id        int         `json:"id"`
	Type      string      `json:"type"`
	UpdatedAt time.Time   `json:"updated_at"`
	CreatedAt time.Time   `json:"created_at"`
	Tasks     []TaskDatas `json:"Tasks"`
}

type GetResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

// ======= PUT ========
type UpdateRequest struct {
	Type string `json:"type" valid:"required~type cannot be empty"`
}

type UpdateResponse struct {
	Id                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type UpdateCategoryResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
