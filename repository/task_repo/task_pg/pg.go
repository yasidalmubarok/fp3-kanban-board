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
			user_id,
			title, 
			description,
			category_id
		)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id, title, description, status, user_id, category_id, created_at;
	`

	getTaskById = `
		SELECT 
			t.id,
			t.title,
			t.description,
			t.status,
			t.user_id,
			t.category_id,
			t.created_at,
			u.id,
			u.email,
			u.full_name
		FROM 
			tasks AS t
		LEFT JOIN
			users AS u
		ON
			t.user_id = u.id
		WHERE t.id = $1
	`


)
// LEFT JOIN
// 			categories AS c
// 		ON
// 			t.category_id = c.id
type taskPG struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) task_repo.Repository{
	return &taskPG{
		db: db,
	}
}

func (t *taskPG) CreateNewTask(taskPayLoad *entity.Task) (*dto.NewTasksResponse, errs.MessageErr) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong "+ err.Error())
	}

	var task dto.NewTasksResponse

	row := tx.QueryRow(
		createTask,
		taskPayLoad.UserId,
		taskPayLoad.Title, 
		taskPayLoad.Description, 
		taskPayLoad.CategoryId,
	)
	err = row.Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.UserId,
		&task.CategoryId,
		&task.CreatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return &task, nil
}

func (t *taskPG) GetTaskById(id int) (*task_repo.TaskUserMapped, errs.MessageErr) {
	var taskUser task_repo.TaskUser
	
	err := t.db.QueryRow(getTaskById, id).Scan(
		&taskUser.Task.Id,
		&taskUser.Task.Title,
		&taskUser.Task.Description,
		&taskUser.Task.Status,
		&taskUser.Task.UserId,
		&taskUser.Task.CategoryId,
		&taskUser.Task.CreatedAt,
		&taskUser.User.Id,
		&taskUser.User.Email,
		&taskUser.User.FullName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("task not found" + err.Error())
		}
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	result := task_repo.TaskUserMapped{}
	return result.HandleMappingTaskUser(taskUser), nil
}