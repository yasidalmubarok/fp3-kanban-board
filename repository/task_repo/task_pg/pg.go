package task_pg

import (
	"database/sql"
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/repository/task_repo"
)

const(
	createTask = `
		INSERT INTO tasks (
			title, 
			description,
			category_id,
		)
		VALUES ($1, $2, $3)
		RETURNING
			id, title, status, description, user_id, category_id, created_at;
	`
)

type taskPG struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) task_repo.Repository{
	return &taskPG{
		db: db,
	}
}

func (t taskPG) CreateNewTask(taskPayLoad *entity.Task) (*dto.NewTasksResponse, errs.MessageErr) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	var task dto.NewTasksResponse

	row := tx.QueryRow(
		createTask,
		taskPayLoad.Title, 
		taskPayLoad.Description, 
		taskPayLoad.CategoryId,
	)
	err = row.Scan(
		&task.Id,
		&task.Title,
		&task.Status,
		&task.Description,
		&task.UserId,
		&task.CategoryId,
		&task.CreatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &task, nil
}