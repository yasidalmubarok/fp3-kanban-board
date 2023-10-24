package task_repo

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
)

type Repository interface {
	CreateNewTask(taskPayLoad *entity.Task) (*dto.NewTasksResponse, errs.MessageErr)
	GetTask() ([]TaskUserMapped, errs.MessageErr)
	GetTaskById(id int) (*TaskUserMapped, errs.MessageErr)
}
