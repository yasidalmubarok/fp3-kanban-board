package category_service

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/pkg/helper"
	"final-project/repository/category_repo"
	"final-project/repository/task_repo"
	"net/http"
)

type CategoryService interface {
	Create(categoryPayLoad *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr)
	Get() (*dto.GetResponse, errs.MessageErr)
}

type categoryService struct {
	categoryRepo category_repo.Repository
	taskRepo     task_repo.Repository
}

// factory function
func NewCategorySevice(categoryRepo category_repo.Repository, taskRepo task_repo.Repository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
		taskRepo: taskRepo,
	}
}

// Implements service interface
func (cs *categoryService) Create(categoryPayLoad *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr) {
	err := helper.ValidateStruct(categoryPayLoad)

	if err != nil {
		return nil, err
	}

	category := &entity.Category{
		Type: categoryPayLoad.Type,
	}

	response, err := cs.categoryRepo.Create(category)

	if err != nil {
		return nil, err
	}

	response = &dto.NewCategoryResponse{
		Id:        response.Id,
		Type:      response.Type,
		CreatedAt: response.CreatedAt,
	}

	return response, nil
}

func (cs *categoryService) Get() (*dto.GetResponse, errs.MessageErr) {
	categories, err := cs.categoryRepo.Read()

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	response := dto.GetResponse{
		StatusCode: http.StatusOK,
		Message:    "categories successfully fetched",
		Data:       categories,
	}

	return &response, nil
}
