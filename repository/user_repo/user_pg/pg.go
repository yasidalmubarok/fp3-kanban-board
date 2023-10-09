package user_pg

import (
	"database/sql"
	"errors"
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/repository/user_repo"

	_ "github.com/lib/pq"
)

const (
	getUserByEmail = `
		SELECT id, email, password FROM "users"
		WHERE email = $1;
	`

	getUserById = `
		SELECT id, email, password FROM "users"
		WHERE id = $1;
	`

	createNewUser = `
		INSERT INTO "users"
		(
			full_name,
			email,
			password,
			role
		)
		VALUES ($1, $2, $3, 'member')
		RETURNING
		id, full_name, email, created_at
	`


	updateUser = `
		UPDATE "users"
		SET "full_name" = $2,
		"email" = $3
		WHERE "id" = $1
		RETURNING
			id, full_name, email, updated_at
	`
)

type userPG struct {
	db *sql.DB
}

func NewUserPG(db *sql.DB) user_repo.Repository {
	return &userPG{
		db: db,
	}
}

// Membuat implementasi interface user_repo
func (u *userPG) CreateNewUser(userPayLoad *entity.User) (*dto.NewUserResponse, errs.MessageErr) {
	tx, err := u.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	var user dto.NewUserResponse
	row := tx.QueryRow(createNewUser, userPayLoad.FullName, userPayLoad.Email, userPayLoad.Password)

	err = row.Scan(&user.Id, &user.FullName, &user.Email, &user.CreatedAt)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError(err.Error())
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (u *userPG) UpdateUser(oldUser *entity.User, newUser *entity.User) (*entity.User, errs.MessageErr) {
	tx, err := u.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(updateUser, newUser.FullName, newUser.Email, oldUser.Id)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return oldUser, nil
}

func (u *userPG) GetUserByEmail(email string) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(getUserByEmail, email)

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (u *userPG) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	var user entity.User

	row := u.db.QueryRow(getUserByEmail, userId)

	err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}
