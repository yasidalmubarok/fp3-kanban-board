package dto

import "time"


type TaskDatas struct {
	Id          uint       `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserId      uint       `json:"user_id"`
	CategoryId  uint       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}
type NewTasksRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryId  uint    `json:"category_id"`
}

type NewTasksResponse struct {
	Id          uint       `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserId      uint       `json:"user_id"`
	CategoryId  uint       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}
