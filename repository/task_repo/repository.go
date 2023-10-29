package task_repo

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
)

type Repository interface {
	CreateNewTask(taskPayLoad *entity.Task) (*dto.NewTasksResponse, errs.MessageErr)
	GetTask() ([]TaskUserMapped, errs.MessageErr)
	GetTaskById(id int) (*entity.Task, errs.MessageErr)
	UpdateTaskById(taskPayLoad *entity.Task) (*dto.UpdateTaskResponse, errs.MessageErr)
	UpdateTaskByStatus(taskPayLoad *entity.Task) (*dto.UpdateTaskResponseByStatus, errs.MessageErr)
	UpdateTaskByCategoryId(taskPayLoad *entity.Task) (*dto.UpdateCategoryIdResponse, errs.MessageErr)
	DeleteTaskById(taskId int) (errs.MessageErr)
}
