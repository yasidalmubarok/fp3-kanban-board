package category_repo

import "final-project/entity"

type CategoryTaskMapped struct {
	Category entity.Category
	Tasks []entity.Task
}

type CategoryTask struct {
	Category entity.Category
	Task entity.Task
}

func (ctm *CategoryTaskMapped) HandleMappingCategoryWithTask(categoryTask []CategoryTask) []CategoryTaskMapped {
	categoryTasksMapped := []CategoryTaskMapped{}

	for _, eachCategoryTask := range categoryTask{
		isCategoryExist := false

		for i := range categoryTasksMapped{
			if eachCategoryTask.Category.Id == categoryTasksMapped[i].Category.Id{
				isCategoryExist = true
				categoryTasksMapped[i].Tasks = append(categoryTasksMapped[i].Tasks, eachCategoryTask.Task)
				break
			}
		}

		if !isCategoryExist{
			categoryTaskMapped := CategoryTaskMapped{
				Category: eachCategoryTask.Category,
			}

			categoryTaskMapped.Tasks = append(categoryTaskMapped.Tasks, eachCategoryTask.Task)

			categoryTasksMapped = append(categoryTasksMapped, categoryTaskMapped)  
		}
	}

	return categoryTasksMapped
}