package category_pg

import (
	"database/sql"
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
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
)

type categoryPG struct {
	db *sql.DB
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

	if err := tx.Commit(); err != nil{
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &category, nil
}
