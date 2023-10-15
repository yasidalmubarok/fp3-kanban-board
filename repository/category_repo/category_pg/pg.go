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
			"c"."id", 
			"c"."type", 
			"c"."updated_at", 
			"c"."created_at",
			"t"."id", 
			"t"."title", 
			"t"."description", 
			"t"."user_id", 
			"t"."category_id", 
			"t"."created_at", 
			"t"."updated_at"
		FROM 
			categories AS "c"
		LEFT JOIN 
			"tasks" AS "t" 
		ON 
			"c"."id" = "t"."category_id"
		ORDER BY 
			"c"."id" 
		ASC
	`

	checkCategoryId = `
		SELECT 
			c.id 
		FROM 
			categories AS c
		WHERE
			c.id = $1
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

func (c categoryPG) Read() ([]category_repo.CategoryTaskMapped, errs.MessageErr) {
	categoryTasks := []category_repo.CategoryTask{}

	rows, err := c.db.Query(getCategoryWithTask)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong" + err.Error())
	}

	for rows.Next() {
		var categoryTask category_repo.CategoryTask

		err := rows.Scan(
			&categoryTask.Category.Id,
			&categoryTask.Category.Type,
			&categoryTask.Category.CreatedAt,
			&categoryTask.Category.UpdatedAt,

			&categoryTask.Task.Id,
			&categoryTask.Task.Title,
			&categoryTask.Task.Description,
			&categoryTask.Task.UserId,
			&categoryTask.Task.CategoryId,
			&categoryTask.Task.CreatedAt,
			&categoryTask.Task.UpdatedAt,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong" + err.Error())
		}

		categoryTasks = append(categoryTasks, categoryTask)
	}

	var result category_repo.CategoryTaskMapped

	return result.HandleMappingCategoryWithTask(categoryTasks), nil
}

func (c *categoryPG) ReadById(categoryId int) (*category_repo.CategoryTaskMapped, errs.MessageErr) {
	category := category_repo.CategoryTask{}
	row := c.db.QueryRow(checkCategoryId, categoryId)
	err := row.Scan(
		&category.Category.Id,
		// &category.Category.Type,
		// &category.Category.CreatedAt,
		// &category.Category.UpdatedAt,

		// &category.Task.Id,
		// &category.Task.Title,
		// &category.Task.Description,
		// &category.Task.UserId,
		// &category.Task.CategoryId,
		// &category.Task.CreatedAt,
		// &category.Task.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewInternalServerError("rows not found" + err.Error())
		}
		return nil, errs.NewInternalServerError("something went wrong " + err.Error())
	}

	result := category_repo.CategoryTaskMapped{}
	return result.HandleMappingCategoryWithTaskById(category), nil
}
