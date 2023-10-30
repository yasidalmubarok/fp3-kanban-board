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

// Create implements TasksHandler.
// Create godoc
// @Summary Create new Task
// @Description Create new Task
// @Tags Task
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param RequestBody body dto.NewTasksRequest true "body request for add new Task"
// @Success 201 {object} dto.NewTasksResponse
// @Router /tasks [post]
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

// Get implements TasksHandler.
// Get godoc
// @Summary Get Tasks
// @Description Get Tasks
// @Tags Task
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} dto.GetResponseTasks
// @Router /tasks [get]
func (th *taskHandler) Get(ctx *gin.Context) {
	response, err := th.taskService.Get()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// Update implements TaskHandler.
// Update godoc
// @Summary Update task
// @Description Update task
// @Tags Task
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param taskId path int true "taskId"
// @Param RequestBody body dto.UpdateTaskRequest true "body request for update task"
// @Success 200 {object} dto.UpdateResponseTask
// @Router /tasks/{taskId} [put]
func (th *taskHandler) Update(ctx *gin.Context) {
	taskId, _ := strconv.Atoi(ctx.Param("taskId"))

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

	ctx.JSON(response.StatusCode, response)
}

// UpdateByStatus implements TaskHandler.
// UpdateByStatus godoc
// @Summary Update task by status
// @Description Update task by status
// @Tags Task
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param taskId path int true "taskId"
// @Param RequestBody body dto.UpdateTaskRequestByStatus true "body request for update task by status"
// @Success 200 {object} dto.UpdateResponseTask
// @Router /tasks/update-status/{taskId} [patch]
func (th *taskHandler) UpdateByStatus(ctx *gin.Context) {
	taskId, _ := strconv.Atoi(ctx.Param("taskId"))

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

	ctx.JSON(response.StatusCode, response)
}

// UpdateBycategoryId implements TaskHandler.
// UpdateBycategoryId godoc
// @Summary Update task by categoryId
// @Description Update task by categoryId
// @Tags Task
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param taskId path int true "taskId"
// @Param RequestBody body dto.UpdateCategoryIdRequest true "body request for update task by status"
// @Success 200 {object} dto.UpdateCategoryId
// @Router /tasks/update-category/{taskId} [patch]
func (th *taskHandler) UpdateByCategoryId(ctx *gin.Context) {
	taskId, _ := strconv.Atoi(ctx.Param("taskId"))

	task := &dto.UpdateCategoryIdRequest{}

	if err := ctx.ShouldBindJSON(task); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := th.taskService.UpdateTaskByCategoryId(taskId, task)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// Delete implements TasksHandler.
// Delete godoc
// @Summary Delete
// @Description Delete
// @Tags Task
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param taskId path int true "taskId"
// @Success 200 {object} dto.DeleteTaskByIdResponse
// @Router /tasks/{taskId} [delete]
func (th *taskHandler) Delete(ctx *gin.Context) {
	taskId, _ := strconv.Atoi(ctx.Param("taskId"))

	response, err := th.taskService.DeleteTaskById(taskId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
