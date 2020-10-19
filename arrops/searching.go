package arrops

// LinearSearch uses linear search algorithm to obtain the given index of the value,
// Returns -1 if not found
func (arr *ArrayADT) LinearSearch(value int) int {
	index := -1
	for i := 0; i < arr.length; i++ {
		if arr.data[i] == value {
			index = i
			break
		}
	}
	return index
}
