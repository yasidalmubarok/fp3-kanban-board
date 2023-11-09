package category_pg

import (
	"database/sql"
	"final-project/entity"
	"final-project/repository/category_repo"
	"time"
)

type categoryWithTask struct {
	CategoryId        int
	CategoryType      string
	CategoryCreatedAt time.Time
	CategoryUpdatedAt time.Time
	TaskId            sql.NullInt64
	TaskTitle         sql.NullString
	TaskDescription   sql.NullString
	TaskStatus        sql.NullBool
	TaskUserId        sql.NullInt64
	TaskCategoryId    sql.NullInt64
	TaskCreatedAt     sql.NullTime
	TaskUpdatedAt     sql.NullTime
}

func (c *categoryWithTask) categoryWithTaskToEntity() category_repo.CategoryTask {
	return category_repo.CategoryTask{
		Category: entity.Category{
			Id:        c.CategoryId,
			Type:      c.CategoryType,
			CreatedAt: c.CategoryCreatedAt,
			UpdatedAt: c.CategoryUpdatedAt,
		},
		Task: entity.Task{
			Id:          int(c.TaskId.Int64),
			Title:       c.TaskTitle.String,
			Description: c.TaskDescription.String,
			Status:      c.TaskStatus.Bool,
			UserId:      int(c.TaskUserId.Int64),
			CategoryId:  int(c.TaskCategoryId.Int64),
			CreatedAt:   c.TaskCreatedAt.Time,
			UpdatedAt:   c.TaskUpdatedAt.Time,
		},
	}
}
