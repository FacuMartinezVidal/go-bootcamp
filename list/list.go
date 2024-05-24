package list

func List() {
	//array, but internally it's a slice
	prices := []int{1, 2, 3, 4, 5, 7}
	//the slice is from the same array, don't use more space in memory
	slice := prices[:3]
	//capacity takes the length of the array, not the slice
	cap := cap(prices)
	//len takes the length of the slice, not the array
	len := len(prices)
	//append a new element to the slice
	prices = append(prices, 8)
	//go has't a method to remove an element from a slice, we can create a new slice without the element
	newSlice := slice[1:]

	//for append a slice to another slice, we have to use the ... operator, to unpack the slice
	appendSlice := append(prices, slice...)
}
