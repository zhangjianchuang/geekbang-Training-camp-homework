package v1

import "github.com/google/wire"

type Name string
type User struct {
	N Name
}

func NewUser(name Name) User {
	return User{N: name}
}

// Create add new user to the storage.
func Create(n Name) User {
	wire.Build(NewUser)
	return User{}
}
