package category_repo

import (
	"final-project/entity"
	"time"
)

type CategoryTask struct {
	Category entity.Category
	Task     entity.Task
}

type CategoryTaskMapped struct {
	Id        int           `json:"id"`
	Type      string        `json:"type"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Tasks     []entity.Task `json:"Task"`
}

func (ctm *CategoryTaskMapped) HandleMappingCategoryWithTask(categoryTask []*CategoryTask) []*CategoryTaskMapped {
	categoryTasksMapped := []*CategoryTaskMapped{}

	for _, eachCategoryTask := range categoryTask {
		isCategoryExist := false

		for i := range categoryTasksMapped {
			if eachCategoryTask.Category.Id == categoryTask[i].Category.Id {
				isCategoryExist = true
				categoryTasksMapped[i].Tasks = append(categoryTasksMapped[i].Tasks, eachCategoryTask.Task)
				break
			}
		}

		if !isCategoryExist {
			categoryTaskMapped := &CategoryTaskMapped{
				Id:        eachCategoryTask.Category.Id,
				Type:      eachCategoryTask.Category.Type,
				CreatedAt: eachCategoryTask.Category.CreatedAt,
				UpdatedAt: eachCategoryTask.Category.UpdatedAt,
			}

			categoryTaskMapped.Tasks = append(categoryTaskMapped.Tasks, eachCategoryTask.Task)			
			categoryTasksMapped = append(categoryTasksMapped, categoryTaskMapped)

			if categoryTaskMapped.Tasks[0].Id == 0 {
				categoryTaskMapped.Tasks = []entity.Task{}
			}
		}
	}

	return categoryTasksMapped
}
