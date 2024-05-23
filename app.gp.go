package main

import (
	"fmt"
	"go-bootcamp/user"
)

type stringer interface {
	String() string
}

// declaring function
func main() {

	u, err := user.New("John", 25, "jhon@gmail")
	if err != nil {
		fmt.Println(err)
		return
	}

	a, err := user.NewAdmin("John", 25, "jhon@gmail", "admin", "password")
	if err != nil {
		fmt.Println(err)
		return
	}

	printer(u)
	printer(a)

}

func printer(s stringer) {
	fmt.Println(s.String())
}
