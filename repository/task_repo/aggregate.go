package task_repo

import (
	"final-project/entity"
	"time"
)

type user struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type TaskUser struct {
	Task entity.Task
	User entity.User
}

type TaskUserMapped struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int       `json:"userId"`
	CategoryId  int       `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
	User        user      `json:"user"`
}

func (tum *TaskUserMapped) HandleMappingTasksUser(taskUser []TaskUser) []TaskUserMapped {
	tasksUserMapped := []TaskUserMapped{}

	for _, eachTaskUser := range taskUser {
		taskUserMapped := TaskUserMapped{
			Id:          eachTaskUser.Task.Id,
			Title:       eachTaskUser.Task.Title,
			Description: eachTaskUser.Task.Description,
			Status:      eachTaskUser.Task.Status,
			UserId:      eachTaskUser.Task.UserId,
			CategoryId:  eachTaskUser.Task.CategoryId,
			CreatedAt:   eachTaskUser.Task.CreatedAt,
			User: user{
				Id:       eachTaskUser.User.Id,
				Email:    eachTaskUser.User.Email,
				FullName: eachTaskUser.User.FullName,
			},
		}
		tasksUserMapped = append(tasksUserMapped, taskUserMapped)
	}
	return tasksUserMapped
}

func (tum *TaskUserMapped) HandleMappingTaskUser(taskUser TaskUser) *TaskUserMapped {
	return &TaskUserMapped{
		Id:          taskUser.Task.Id,
		Title:       taskUser.Task.Title,
		Description: taskUser.Task.Description,
		Status:      taskUser.Task.Status,
		UserId:      taskUser.Task.UserId,
		CategoryId:  taskUser.Task.CategoryId,
		CreatedAt:   taskUser.Task.CreatedAt,
		User: user{
			Id:       taskUser.User.Id,
			Email:    taskUser.User.Email,
			FullName: taskUser.User.FullName,
		},
	}
}
