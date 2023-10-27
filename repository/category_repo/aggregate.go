package category_repo

import (
	"final-project/entity"
	"time"
)

type task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"userId"`
	CategoryId  int       `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CategoryTask struct {
	Category entity.Category
	Task     entity.Task
}

type CategoryTaskMapped struct {
	Id        int       `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tasks     []task    `json:"Task"`
}

func (ctm *CategoryTaskMapped) HandleMappingCategoryWithTask(categoryTask []CategoryTask) []CategoryTaskMapped {
	categoryTasksMap := make(map[int]CategoryTaskMapped)

	for _, eachCategoryTask := range categoryTask {
		categoryId := eachCategoryTask.Category.Id
		categoryTaskMap, exists := categoryTasksMap[categoryId]
		if !exists {
			categoryTaskMap = CategoryTaskMapped{
				Id:        eachCategoryTask.Category.Id,
				Type:      eachCategoryTask.Category.Type,
				CreatedAt: eachCategoryTask.Category.CreatedAt,
				UpdatedAt: eachCategoryTask.Category.UpdatedAt,
			}
		}

		taskItem := task{
			Id:          eachCategoryTask.Task.Id,
			Title:       eachCategoryTask.Task.Title,
			Description: eachCategoryTask.Task.Description,
			UserId:      eachCategoryTask.Task.UserId,
			CategoryId:  eachCategoryTask.Task.CategoryId,
			CreatedAt:   eachCategoryTask.Task.CreatedAt,
			UpdatedAt:   eachCategoryTask.Task.UpdatedAt,
		}
		categoryTaskMap.Tasks = append(categoryTaskMap.Tasks, taskItem)
		categoryTasksMap[categoryId] = categoryTaskMap
	}

	categoryTasks := []CategoryTaskMapped{}
	for _, categoryTask := range categoryTasksMap {
		categoryTasks = append(categoryTasks, categoryTask)
	}

	return categoryTasks
}
