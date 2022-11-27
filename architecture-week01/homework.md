我觉得遇到sql.ErrNoRows的时候应该分两种情况

1. 这个error就不是一个error，没有数据是一个正常现象，比如登陆功能，很多都是如果发现用户不存在，直接就给他注册了，而不是向页面说这个用户不存在

2. 这个error在业务上也是一个error，那么就应该Wrap，比如说我要获取一个用户和它的年龄信息，这个信息需要查两次数据库，如果不Wrap，那么上层根本不知道到底用户不存在，还是年龄不存在。另外我认为哪怕只是查一次数据库，也应该去封装，这样能够精确定位到问题在哪里

```go
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

// 如果用户不存在，去创建一个，所以ErrNoRows不是异常
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
```
