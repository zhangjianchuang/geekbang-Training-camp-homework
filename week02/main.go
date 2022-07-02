package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	age int
}

func getUserFromDB(username string) (User, error) {
	return User{}, sql.ErrNoRows
}
func getUserAgeFromDB(username string) (int, error) {
	return 0, sql.ErrNoRows
}

func addUser() User {
	return User{}
}

func getUser(username string) (User, error) {
	u, err := getUserFromDB(username)
	if err != nil && err == sql.ErrNoRows {
		return addUser(), nil
	}
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func getUserWithAge(username string) (User, error) {
	u, err := getUserFromDB(username)
	if err != nil && err == sql.ErrNoRows {
		return User{}, errors.Wrap(sql.ErrNoRows, fmt.Sprintf("user[%s] not found", username))
	}
	a, err := getUserAgeFromDB(username)
	if err != nil && err == sql.ErrNoRows {
		return User{}, errors.Wrap(sql.ErrNoRows, fmt.Sprintf("user[%s] age not found", username))
	}
	if err != nil {
		return User{}, err
	}
	u.age = a
	return u, nil
}

func main() {
	_, err := getUser("abc")
	if err != nil {
		fmt.Println("getUser got an error: ", err.Error())
	}
	_, err = getUserWithAge("abc")
	if err != nil {
		fmt.Println("getUserWithAge got an error: ", err.Error())
	}
}
