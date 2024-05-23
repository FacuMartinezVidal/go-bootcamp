package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	name      string
	age       int
	email     string
	createdAt time.Time
}

type Admin struct {
	User
	credential string
	password   string
}

// constructor

func NewAdmin(name string, age int, email string, credential string, password string) (*Admin, error) {
	if name == "" && age < 0 && email == "" {
		return nil, errors.New("Invalid input")
	}
	return &Admin{
		User: User{
			name:  name,
			age:   age,
			email: email,
		},
		credential: credential,
		password:   password,
	}, nil

}

func New(name string, age int, email string) (*User, error) {
	if name == "" && age < 0 && email == "" {
		return nil, errors.New("Invalid input")
	}
	return &User{
		name:  name,
		age:   age,
		email: email,
	}, nil
}

func (u *User) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Email: %s", u.name, u.age, u.email)
}

// method
// we have to pass the receiver as a pointer to the struct, so that we can modify the struct\

func (a *Admin) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Email: %s, Credential: %s, Password: %s", a.name, a.age, a.email, a.credential, a.password)
}
