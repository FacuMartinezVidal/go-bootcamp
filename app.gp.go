package main

import (
	"fmt"
)

type stringer interface {
	String() string
}

// declaring function
func main() {
	
}

func printer(s stringer) {
	fmt.Println(s.String())
}
