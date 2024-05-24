package make

func Make() {
	// make() is used to create a slice, map, or channel
	// with make() you can help the compiler to allocate the correct amount of memory
	array := make([]int, 5)

	// make() with 3 arguments
	// make(type, length, capacity)
	// capacity is optional
	array2 := make([]int, 5, 10)

	// make() with map
	map1 := make(map[string]int)
	
}
