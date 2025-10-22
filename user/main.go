package user

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (fn User) GetFullName() string {
	return fmt.Sprintf("%v %v", fn.FirstName, fn.LastName)
}

var Users []User

var UserLogin User
