package main

import "fmt"

// declaring function
func main() {
	// Declaring variables
	//var a int
	//var b = 60

	//taking input from user, & is used to get the address of the variable, it called pointer
	//fmt.Scan(&a)

	//declaring multiple variables
	d, e := 14, 20

	//declaring constant
	const f = 10

	add(d, e)

	const text = "Hello, World!"
	output(text)

	//if statement
	var h, i int
	fmt.Scan(&h)
	fmt.Scan(&i)
	if h == 10 && i == 10 {
		println("Both are equal")
	} else {
		println("Both are not equal")
	}

	//for loop
	for j := 0; j < 5; j++ {
		println(j)
		if j == 3 {
			break
		} else {
			continue
		}

	}

	//switch case
	var k int
	fmt.Scan(&k)
	switch k {
	case 1:
		println("One")
	case 2:
		println("Two")
	case 3:
		println("Three")
	}

}

func output(text string) {
	println(text)
}

func add(a int, b int) (int, string) {
	return a + b, "Addition is done"
}

func sub(a int, b int) (result int) {
	result = a - b
	//return result
	return
}
