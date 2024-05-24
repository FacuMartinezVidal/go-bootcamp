package _for

func For() {
	// declaring for loop
	for i := 0; i < 5; i++ {
		// printing value of i
		println(i)
	}

	// range loop, you can use range to iterate over an array, slice, string, map, or channel
	s := []int{1, 2, 3}
	for index, value := range s {
		println(index, value)
	}

}
