package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	age int
}

func getUserFromDB(_ string) (User, error) {
	return User{}, sql.ErrNoRows
}
func getUserAgeFromDB(_ string) (int, error) {
	return 0, sql.ErrNoRows
}

func addUser(_ string) User {
	return User{}
}

// 如果用户不存在，去创建一个，所以ErrNoRows不是异常
func getUser(username string) (User, error) {
	u, err := getUserFromDB(username)
	if err != nil && err == sql.ErrNoRows {
		return addUser(username), nil
	}
	if err != nil {
		return User{}, err
	}
	return u, nil
}

// 用户和年龄都必须存在，那么有一个不存在都需要向上返回一个错误
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
		fmt.Println(fmt.Errorf("getUser got an error:%+v ", err))
	}
	_, err = getUserWithAge("abc")
	if err != nil {
		fmt.Println(fmt.Errorf("getUserWithAge got an error: %+v", err))
	}
}
