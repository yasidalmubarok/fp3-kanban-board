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

//====== POST ======
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

//======= GET ========
type GetTaskResponse struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Status      bool       `json:"status"`
	Description string     `json:"description"`
	UserId      int        `json:"user_id"`
	CategoryId  int        `json:"category_id"`
	CreatedAt   time.Time  `json:"created_at"`
	Users       []GetUsers `json:"Users"`
}

type GetResponseTasks struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

// ======= UPDATE ========

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type UpdateTaskRequestByStatus struct {
	Status bool `json:"status"`
}
type UpdateTaskResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type UpdateTaskResponseByStatus struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateResponseTask struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type UpdateCategoryIdRequest struct {
	CategoryId int `json:"category_id"`
}

type UpdateCategoryIdResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateCategoryId struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

// ========= DELETE ==========
type DeleteTaskByIdResponse struct {
	Message string `json:"message"`
}