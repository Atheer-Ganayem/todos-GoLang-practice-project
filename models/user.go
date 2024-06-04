package models

import (
	"errors"

	db "todo-golang.com/DB"
	"todo-golang.com/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) Save() error {
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password cannot be empty")
	}
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {return err}
	defer stmt.Close()

	hashedPw, err := utils.HashPassword(user.Password)
	if err != nil {return err}

	result, err := stmt.Exec(user.Email, hashedPw)
	if err != nil {return err}

	id, err := result.LastInsertId()
	if err != nil {return err}

	user.ID = id

	return nil
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users where email = ?"
	row := db.DB.QueryRow(query, user.Email)
	
	var retrievedPassword string
	if err := row.Scan(&user.ID, &retrievedPassword); err != nil {
		return err
	}

	isPasswordValid := utils.CompareStringToHash(user.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("invalid email or password")
	}

	return nil
}