package arrops

// LinearSearch uses linear search algorithm to obtain the given index of the value
// Returns -1 if not found
func (arr *ArrayADT) LinearSearch(value int) int {
	for i := 0; i < arr.length; i++ {
		if arr.data[i] == value {
			return i
		}
	}
	return -1
}

// Contains returns true if the value is present on the array and returns false otherwise
func (arr *ArrayADT) Contains(value int) bool {
	if arr.LinearSearch(value) != -1 {
		return true
	}
	return false
}
