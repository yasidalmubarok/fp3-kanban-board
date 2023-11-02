package category_handler

import (
	"final-project/dto"
	"final-project/pkg/errs"
	"final-project/service/category_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService category_service.CategoryService
}

func NewCategoryHandler(categoryService category_service.CategoryService) *categoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}

// Create implements CategoriesHandler.
// Create godoc
// @Summary Create new Category
// @Description Create new Category
// @Tags Category
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param RequestBody body dto.NewCategoryRequest true "body request for add new Task"
// @Success 201 {object} dto.NewCategoryResponse
// @Router /categories [post]
func (ch *categoryHandler) Create(ctx *gin.Context) {
	var newCategoryRequest = &dto.NewCategoryRequest{}

	if err := ctx.ShouldBindJSON(newCategoryRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := ch.categoryService.Create(newCategoryRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// Get implements CategoriesHandler.
// Get godoc
// @Summary Get Tasks
// @Description Get Categories
// @Tags Category
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} dto.GetResponse
// @Router /categories [get]
func (ch *categoryHandler) Get(ctx *gin.Context) {
	response, err := ch.categoryService.Get()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}

// Update implements CategoriesHandler.
// Update godoc
// @Summary Update Category
// @Description Update Category
// @Tags Category
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param categoryId path int true "categoryId"
// @Param RequestBody body dto.UpdateRequest true "body request for update task"
// @Success 200 {object} dto.UpdateCategoryResponse
// @Router /categories/{categoryId} [patch]
func (ch *categoryHandler) Update(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Param("categoryId"))

	category := &dto.UpdateRequest{}

	if err := ctx.ShouldBindJSON(category); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := ch.categoryService.Update(categoryId, category)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Delete implements CategoriesHandler.
// Delete godoc
// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param categoryId path int true "categoryId"
// @Success 200 {object} dto.DeleteCategoryByIdResponse
// @Router /categories/{categoryId} [delete]
func (ch *categoryHandler) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Param("categoryId"))

	response, err := ch.categoryService.Delete(categoryId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
