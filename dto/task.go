package dto

import "time"

type TaskDatas struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type NewTasksRequest struct {
	Title       string `json:"title" valid:"required~full_name cannot be empty"`
	Description string `json:"description" valid:"required~full_name cannot be empty"`
	CategoryId  int    `json:"category_id"`
}

type NewTasksResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}
