package category_repo

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
)

type Repository interface {
	Create(categoryPayLoad *entity.Category) (*dto.NewCategoryResponse, errs.MessageErr)
	Read() ([]CategoryTaskMapped, errs.MessageErr)
	ReadById(id uint) (*CategoryTaskMapped, errs.MessageErr)
}