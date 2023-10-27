package taks_handler

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/service/task_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type taskHandler struct {
	taskService task_service.TaskService
}

func NewTaskHandler(taskService task_service.TaskService) *taskHandler {
	return &taskHandler{
		taskService: taskService,
	}
}

func (th *taskHandler) Create(ctx *gin.Context) {
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

func (th *taskHandler) Get(ctx *gin.Context) {
	response, err := th.taskService.Get()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (th *taskHandler) Update(ctx *gin.Context) {
	taskId, _:= strconv.Atoi(ctx.Param("taskId"))

	task := &dto.UpdateTaskRequest{}

	if err := ctx.ShouldBindJSON(task); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := th.taskService.UpdateTask(taskId, task)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (th *taskHandler) UpdateByStatus(ctx *gin.Context) {
	taskId, _:= strconv.Atoi(ctx.Param("taskId"))

	task := &dto.UpdateTaskRequestByStatus{}

	if err := ctx.ShouldBindJSON(task); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := th.taskService.UpdateTaskByStatus(taskId, task)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
