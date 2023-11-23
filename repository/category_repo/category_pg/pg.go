package category_pg

import (
	"database/sql"
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/repository/category_repo"

	_ "github.com/lib/pq"
)

const (
	createCategory = `
		INSERT INTO "categories"
		(
			type
		)
		VALUES ($1)
		RETURNING
			id, type, created_at;
	`

	getCategoryWithTask = `
		SELECT
			c.id,
			c.type,
			c.updated_at,
			c.created_at,
			t.id,
			t.title,
			t.description,
			t.user_id,
			t.category_id,
			t.created_at,
			t.updated_at
		FROM
			categories AS c
		LEFT JOIN
			tasks AS t
		ON
			c.id = t.category_id
		ORDER BY
			c.id
		ASC
	`

	updateCategoryById = `
		UPDATE
			categories
		SET
			type = $2,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, type, updated_at
	`

	checkCategoryId = `
		SELECT 
			c.id 
		FROM 
			categories AS c
		WHERE
			c.id = $1
	`

	deleteCategoryById = `
		DELETE
		FROM
			categories
		WHERE
			id = $1
	`
)

type categoryPG struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) category_repo.Repository {
	return &categoryPG{
		db: db,
	}
}

func (c *categoryPG) Create(categoryPayLoad *entity.Category) (*dto.NewCategoryResponse, errs.MessageErr) {
	tx, err := c.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	var category dto.NewCategoryResponse

	row := tx.QueryRow(createCategory, categoryPayLoad.Type)
	err = row.Scan(&category.Id, &category.Type, &category.CreatedAt)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &category, nil
}

func (c *categoryPG) GetCategory() ([]*category_repo.CategoryTaskMapped, errs.MessageErr) {
	categoryTasks := []*category_repo.CategoryTask{}
	rows, err := c.db.Query(getCategoryWithTask)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	for rows.Next() {
		categoryTask := categoryWithTask{}

		err := rows.Scan(
			&categoryTask.CategoryId,
			&categoryTask.CategoryType,
			&categoryTask.CategoryCreatedAt,
			&categoryTask.CategoryUpdatedAt,
			&categoryTask.TaskId,
			&categoryTask.TaskTitle,
			&categoryTask.TaskDescription,
			&categoryTask.TaskUserId,
			&categoryTask.TaskCategoryId,
			&categoryTask.TaskCreatedAt,
			&categoryTask.TaskUpdatedAt,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		categoryTasks = append(categoryTasks, categoryTask.categoryWithTaskToEntity())
	}

	result := category_repo.CategoryTaskMapped{}
	return result.HandleMappingCategoryWithTask(categoryTasks), nil
}

func (c *categoryPG) UpdateCategory(categoryPayLoad *entity.Category) (*dto.UpdateResponse, errs.MessageErr) {
	tx, err := c.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	row := tx.QueryRow(updateCategoryById, categoryPayLoad.Id, categoryPayLoad.Type)

	var categoryUpdate dto.UpdateResponse
	err = row.Scan(
		&categoryUpdate.Id,
		&categoryUpdate.Type,
		&categoryUpdate.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &categoryUpdate, nil
}

func (c *categoryPG) CheckCategoryId(categoryId int) (*entity.Category, errs.MessageErr) {
	category := entity.Category{}
	row := c.db.QueryRow(checkCategoryId, categoryId)
	err := row.Scan(&category.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewInternalServerError("category not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &category, nil
}

func (c *categoryPG) DeleteCategory(categoryId int) errs.MessageErr {
	tx, _ := c.db.Begin()

	_, err := tx.Exec(deleteCategoryById, categoryId)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
