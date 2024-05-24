package main

import (
	"fmt"
	"go-bootcamp/list"
)

type stringer interface {
	String() string
}

// declaring function
func main() {
	listed := list.List()
	fmt.Println(*listed)
	sliced := list.Slice()
	fmt.Println(*sliced)

}

func printer(s stringer) {
	fmt.Println(s.String())
}
