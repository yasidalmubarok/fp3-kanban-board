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
	Status      bool      `json:"status"`
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
				Status:      eachCategoryTask.Task.Status,
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

func (ctm *CategoryTaskMapped) HandleMappingCategoryWithTaskById(categoryTask CategoryTask) *CategoryTaskMapped {
	return &CategoryTaskMapped{
		Id: categoryTask.Category.Id,
		// Type:      categoryTask.Category.Type,
		// CreatedAt: categoryTask.Category.CreatedAt,
		// UpdatedAt: categoryTask.Category.UpdatedAt,
		// Tasks: Task{
		// 	Id:          categoryTask.Task.Id,
		// 	Title:       categoryTask.Task.Title,
		// 	Description: categoryTask.Task.Description,
		// 	Status:      categoryTask.Task.Status,
		// 	UserId:      categoryTask.Task.UserId,
		// 	CategoryId:  categoryTask.Task.CategoryId,
		// 	CreatedAt:   categoryTask.Task.CreatedAt,
		// 	UpdatedAt:   categoryTask.Task.UpdatedAt,
		// },
	}
}
