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

func (ch *categoryHandler) Get(ctx *gin.Context) {
	response, err := ch.categoryService.Get()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}

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

func (ch *categoryHandler) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Param("categoryId"))

	response, err := ch.categoryService.Delete(categoryId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
