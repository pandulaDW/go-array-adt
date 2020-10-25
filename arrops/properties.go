package arrops

// GetElement returns the value at a given index
func (arr *ArrayADT) GetElement(index int) int {
	return arr.data[index]
}

// SetElement sets the value at a given index
func (arr *ArrayADT) SetElement(index int, value int) {
	arr.data[index] = value
}

// IsSorted checks if a given function is sorted either ascending or descending
// based on the isAsc argument provided
func (arr *ArrayADT) IsSorted(isAsc bool) bool {

	predicate := func(val1, val2 int) bool {
		if isAsc {
			return val1 < val2
		}
		return val1 > val2
	}

	for i := 0; i < arr.length-1; i++ {
		if !predicate(arr.data[i], arr.data[i+1]) {
			return false
		}
	}

	return true
}
