package taks_handler

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/service/task_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService task_service.TaskService
}

func NewTaskHandler(taskService task_service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (th *TaskHandler) Create(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)
	var newTaskRequest = &dto.NewTasksRequest{}

	if err := ctx.ShouldBindJSON(newTaskRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := th.taskService.Create(user.Id, newTaskRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
