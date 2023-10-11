package entity

import "time"

// 			id SERIAL PRIMARY KEY,
// 			title VARCHAR(255) NOT NULL,
// 			description VARCHAR(255) NOT NULL,
// 			status BOOLEAN NOT NULL,
// 			user_id INT,
// 			category_id INT,
// 			created_at timestamptz DEFAULT NOW(),
// 			updated_at timestamptz DEFAULT NOW()

type Task struct {
	Id          uint       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      uint       `json:"userId"`
	CategoryId  uint       `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
