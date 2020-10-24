package arrops

// Pop the last element from a slice
func (arr *ArrayADT) Pop() {
	arr.data = arr.data[:arr.length-1]
	arr.updateLenCap()
}

// Delete the value at a given index
func (arr *ArrayADT) Delete(index int) {
	if index == arr.length-1 {
		arr.Pop()
		return
	}
	slice1 := arr.data[:index]
	slice2 := arr.data[index+1:]
	arr.data = append(slice1, slice2...)
	arr.updateLenCap()
}
