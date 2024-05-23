package basics

func basics() {
	var h, i int
	if h == 10 && i == 10 {
		println("Both are equal")
	} else {
		println("Both are not equal")
	}
	for j := 0; j < 5; j++ {
		println(j)
		if j == 3 {
			break
		} else {
			continue
		}
	}
	var k int
	switch k {
	case 1:
		println("One")
	case 2:
		println("Two")
	case 3:
		println("Three")
	}
}
func pointer(age *int) int {
	return *age + 10

}

// with upper case function name, it can be accessed from other packages, it is called exported function

func Output(text string) {
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
