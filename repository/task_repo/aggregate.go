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
	Users        []user    `json:"user"`
}

func (tum *TaskUserMapped) HandleMappingTasksUser(taskUser []TaskUser) []TaskUserMapped {
	tasksUserMapped := make(map[int]TaskUserMapped)

	for _, eachTaskUser := range taskUser {
		taskId := eachTaskUser.Task.Id
		taskUserMapped, exists := tasksUserMapped[taskId]
		if !exists {
			taskUserMapped = TaskUserMapped{
				Id:          eachTaskUser.Task.Id,
				Title:       eachTaskUser.Task.Title,
				Description: eachTaskUser.Task.Description,
				Status:      eachTaskUser.Task.Status,
				UserId:      eachTaskUser.Task.UserId,
				CategoryId:  eachTaskUser.Task.CategoryId,
				CreatedAt:   eachTaskUser.Task.CreatedAt,
			}
		}
		user := user{
			Id:       eachTaskUser.User.Id,
			Email:    eachTaskUser.User.Email,
			FullName: eachTaskUser.User.FullName,
		}
		taskUserMapped.Users = append(taskUserMapped.Users, user)
		tasksUserMapped[taskId] = taskUserMapped
	}

	taskUsers := []TaskUserMapped{}
	for _, taskUser := range tasksUserMapped {
		taskUsers = append(taskUsers, taskUser)
	}
	return taskUsers
}

