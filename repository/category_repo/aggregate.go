package category_repo

import (
	"final-project/entity"
	"time"
)

type CategoryTask struct {
	Category entity.Category
	Task     entity.Task
}

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"userId"`
	CategoryId  int       `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
type CategoryTaskMapped struct {
	Id        int       `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tasks     Task      `json:"task"`
}

func (ctm *CategoryTaskMapped) HandleMappingCategoryWithTask(categoryTask []CategoryTask) []CategoryTaskMapped {
	categoryTasksMapped := []CategoryTaskMapped{}

	for _, eachCategoryTask := range categoryTask {
		categoryTaskMapped := CategoryTaskMapped{
			Id:        eachCategoryTask.Category.Id,
			Type:      eachCategoryTask.Category.Type,
			CreatedAt: eachCategoryTask.Category.CreatedAt,
			UpdatedAt: eachCategoryTask.Category.UpdatedAt,
			Tasks: Task{
				Id:          eachCategoryTask.Task.Id,
				Title:       eachCategoryTask.Task.Title,
				Description: eachCategoryTask.Task.Description,
				UserId:      eachCategoryTask.Task.UserId,
				CategoryId:  eachCategoryTask.Task.CategoryId,
				CreatedAt:   eachCategoryTask.Task.CreatedAt,
				UpdatedAt:   eachCategoryTask.Task.UpdatedAt,
			},
		}

		categoryTasksMapped = append(categoryTasksMapped, categoryTaskMapped)
	}

	return categoryTasksMapped
}
