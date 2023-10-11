package entity

import "time"

// id SERIAL PRIMARY KEY,
// type VARCHAR(255) NOT NULL,
// created_at timestamptz DEFAULT NOW(),
// updated_at timestamptz DEFAULT NOW()
type Category struct {
	Id        uint       `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
