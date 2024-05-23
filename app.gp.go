package main

import (
	"fmt"
	"go-bootcamp/user"
)

// declaring function
func main() {

	u, err := user.New("John", 25, "jhon@gmail")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(u)
	}

	a, err := user.NewAdmin("John", 25, "jhon@gmail", "admin", "password")

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(a)
	}

}
