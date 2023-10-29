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

	getTaskWithUser = `
		SELECT
			t.id,
			t.title,
			t.status,
			t.description,
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
		ORDER BY
			t.id
		ASC
	`

	getTaskById = `
		SELECT 
			t.id,
			t.user_id
		FROM 
			tasks AS t
		WHERE 
			t.id = $1
	`

	updateTaskById = `
		UPDATE
			tasks
		SET
			title = $2,
			description = $3,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, title, description, status, user_id, category_id, updated_at
	`

	updateTaskByStatus = `
		UPDATE
			tasks
		SET
			status = $2,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, title, description, status, user_id, category_id, updated_at
	`
	updateTaskByCategoryId = `
		UPDATE
			tasks
		SET
			category_id = $2,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, title, description, status, user_id, category_id, updated_at
	`

	deleteTaskById = `
		DELETE
		FROM
			tasks
		WHERE
			id = $1
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

func (t *taskPG) GetTask() ([]task_repo.TaskUserMapped, errs.MessageErr) {
	tasksUser := []task_repo.TaskUser{}
	rows, err := t.db.Query(getTaskWithUser)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	for rows.Next() {
		var taskUser task_repo.TaskUser

		err := rows.Scan(
			&taskUser.Task.Id,
			&taskUser.Task.Title,
			&taskUser.Task.Status,
			&taskUser.Task.Description,
			&taskUser.Task.UserId,
			&taskUser.Task.CategoryId,
			&taskUser.Task.CreatedAt,
			&taskUser.User.Id,
			&taskUser.User.Email,
			&taskUser.User.FullName,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong " + err.Error())
		}

		tasksUser = append(tasksUser, taskUser)
	}

	result := task_repo.TaskUserMapped{}
	return result.HandleMappingTasksUser(tasksUser), nil
}

func (t *taskPG) GetTaskById(id int) (*entity.Task, errs.MessageErr) {
	var taskUser entity.Task
	
	err := t.db.QueryRow(getTaskById, id).Scan(&taskUser.Id, &taskUser.UserId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("task not found" + err.Error())
		}
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	return &taskUser, nil
}

func (t *taskPG) UpdateTaskById(taskPayLoad *entity.Task) (*dto.UpdateTaskResponse, errs.MessageErr) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	row := tx.QueryRow(updateTaskById, taskPayLoad.Id, taskPayLoad.Title, taskPayLoad.Description)
	
	var taskUpdate dto.UpdateTaskResponse
	err = row.Scan(
		&taskUpdate.Id,
		&taskUpdate.Title,
		&taskUpdate.Description,
		&taskUpdate.Status,
		&taskUpdate.UserId,
		&taskUpdate.CategoryId,
		&taskUpdate.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong "+ err.Error())
	}

	return &taskUpdate, nil
}

func (t *taskPG) UpdateTaskByStatus(taskPayLoad *entity.Task) (*dto.UpdateTaskResponseByStatus, errs.MessageErr) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	row := tx.QueryRow(updateTaskByStatus, taskPayLoad.Id, taskPayLoad.Status)
	
	var taskUpdate dto.UpdateTaskResponseByStatus
	err = row.Scan(
		&taskUpdate.Id,
		&taskUpdate.Title,
		&taskUpdate.Description,
		&taskUpdate.Status,
		&taskUpdate.UserId,
		&taskUpdate.CategoryId,
		&taskUpdate.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong "+ err.Error())
	}

	return &taskUpdate, nil
}

func (t *taskPG) UpdateTaskByCategoryId(taskPayLoad *entity.Task) (*dto.UpdateCategoryIdResponse, errs.MessageErr) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}
	var taskUpdate dto.UpdateCategoryIdResponse
	err = tx.QueryRow(updateTaskByCategoryId, taskPayLoad.Id, taskPayLoad.CategoryId).Scan(
		&taskUpdate.Id,
		&taskUpdate.Title,
		&taskUpdate.Description,
		&taskUpdate.Status,
		&taskUpdate.UserId,
		&taskUpdate.CategoryId,
		&taskUpdate.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &taskUpdate, nil
}

func (t *taskPG) DeleteTaskById(taskId int) (errs.MessageErr) {
	tx, err := t.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(deleteTaskById, taskId)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}