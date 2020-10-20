package arrops

// GetElement returns the value at a given index
func (arr *ArrayADT) GetElement(index int) int {
	return arr.data[index]
}

// SetElement sets the value at a given index
func (arr *ArrayADT) SetElement(index int, value int) {
	arr.data[index] = value
}
