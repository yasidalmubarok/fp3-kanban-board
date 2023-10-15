package database

import (
	"database/sql"
	"final-project/infrastructure/config"
	"final-project/repository/user_repo/user_pg"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	appConfig := config.AppConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DBHost,
		appConfig.DBPort,
		appConfig.DBUser,
		appConfig.DBPassword,
		appConfig.DBName,
	)

	db, err = sql.Open(appConfig.DBDialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while validating database arguments:", err.Error())
		return
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while opening a connection to database:", err.Error())
		return
	}
}

func handleCreateRequiredTables() {
	user_pg.SeedAdmin(db)

	const (
		userTable = `
			CREATE TABLE IF NOT EXISTS "users" (
				id SERIAL PRIMARY KEY,
				full_name VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				password TEXT NOT NULL,
				role TEXT NOT NULL,
				created_at timestamptz DEFAULT now(),
				updated_at timestamptz DEFAULT now(),
				CONSTRAINT 
					unique_email
						UNIQUE(email)
			);
		`
		categoryTable = `
			CREATE TABLE IF NOT EXISTS "categories" (
				id SERIAL PRIMARY KEY,
				type VARCHAR(255) NOT NULL,
				created_at timestamptz DEFAULT now(),
				updated_at timestamptz DEFAULT now()
			);
		`
		taskTable = `
			CREATE TABLE IF NOT EXISTS "tasks" (
				id SERIAL PRIMARY KEY,
				title VARCHAR(255) NOT NULL,
				description VARCHAR(255) NOT NULL,
				status BOOL DEFAULT false,
				user_id INT NOT NULL, 
				category_id INT NOT NULL,
				created_at timestamptz DEFAULT NOW(),
				updated_at timestamptz DEFAULT NOW(),
				CONSTRAINT tasks_user_id_fk
					FOREIGN KEY(user_id)
						REFERENCES users(id)
							ON DELETE CASCADE,
				CONSTRAINT tasks_category_id_fk
					FOREIGN KEY(category_id)
						REFERENCES categories(id)	
							ON DELETE CASCADE					
				);
			`
	)

	_, err = db.Exec(categoryTable)

	if err != nil {
		log.Panic("error while create table categories : ", err.Error())
		return
	}

	_, err = db.Exec(userTable)

	if err != nil {
		log.Panic("error while create table users : ", err.Error())
		return
	}

	_, err = db.Exec(taskTable)

	if err != nil {
		log.Panic("error while create table task : ", err.Error())
		return
	}

	if err != nil {
		log.Panic("error occured while trying to create required tables:", err)
	}
}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
