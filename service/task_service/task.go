package task_service

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/pkg/helper"
	"final-project/repository/category_repo"
	"final-project/repository/task_repo"
	"final-project/repository/user_repo"
)

type TaskService interface {
	Create(userId uint, taskPayLoad *dto.NewTasksRequest) (*dto.NewTasksResponse, errs.MessageErr)
}

type taskService struct {
	taskRepo     task_repo.Repository
	categoryRepo category_repo.Repository
	userRepo     user_repo.Repository
}

func NewTaskService(taskRepo task_repo.Repository, categoryRepo category_repo.Repository, userRepo user_repo.Repository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
		categoryRepo: categoryRepo,
		userRepo: userRepo,
	}
}

func (ts *taskService) Create(userId uint, taskPayLoad *dto.NewTasksRequest) (*dto.NewTasksResponse, errs.MessageErr) {
	err := helper.ValidateStruct(taskPayLoad)

	if err != nil {
		return nil, err
	}

	task := &entity.Task{
		Title:       taskPayLoad.Title,
		Description: taskPayLoad.Description,
		CategoryId:  taskPayLoad.CategoryId,
		UserId:      userId,
	}

	response, err := ts.taskRepo.CreateNewTask(task)

	if err != nil {
		return nil, err
	}

	response = &dto.NewTasksResponse{
		Id:          response.Id,
		Title:       response.Title,
		Status:      response.Status,
		Description: response.Description,
		UserId:      response.UserId,
		CategoryId:  response.CategoryId,
		CreatedAt:   response.CreatedAt,
	}

	return response, nil
}
