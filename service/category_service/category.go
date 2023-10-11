package category_service

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/pkg/helper"
	"final-project/repository/category_repo"
	"net/http"
)

type CategoryService interface {
	Create(categoryPayLoad *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.MessageErr)
	Get() (*dto.GetResponse, errs.MessageErr)
}

type categoryService struct {
	categoryRepo category_repo.Repository
}

// factory function
func NewCategorySevice(categoryRepo category_repo.Repository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
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

	categoryResult := []dto.GetCategoryResponse{}

	for _, eachCategory := range categories {
		category := dto.GetCategoryResponse{
			Id:        uint(eachCategory.Category.Id),
			Type:      eachCategory.Category.Type,
			UpdatedAt: eachCategory.Category.UpdatedAt,
			CreatedAt: eachCategory.Category.CreatedAt,
			Tasks:     []dto.TaskDatas{},
		}

		for _, eachTask := range eachCategory.Tasks{
			task := dto.TaskDatas{
				Id: uint(eachTask.Id),
				Title: eachTask.Title,
				Description: eachTask.Description,
				UserId: uint(eachTask.UserId),
				CategoryId: uint(eachTask.CategoryId),
				CreatedAt: category.CreatedAt,
				UpdatedAt: category.UpdatedAt,
			}
			category.Tasks = append(category.Tasks, task)
		}

		categoryResult = append(categoryResult, category)
	}

	response := dto.GetResponse{
		StatusCode: http.StatusOK,
		Message: "categories successfully fetched",
		Data: categoryResult,
	}

	return &response, nil
}
