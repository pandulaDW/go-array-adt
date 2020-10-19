package arrops

import "math"

// Insert the value at the specified index
func (arr *ArrayADT) Insert(index int, value int) {
	arr.data[index] = value
}

// Push ..., pushes the arguments to the end the of the slice
func (arr *ArrayADT) Push(elements ...int) {
	newLength := arr.length + len(elements)

	if newLength < arr.capacity {
		arr.data = arr.data[:newLength]
		copy(arr.data[arr.length:], elements)
	}

	// making a new slice with extended capacity
	newCap := math.Pow(2, math.Ceil(math.Log2(float64(newLength))))
	newSlice := make([]int, newLength, int(newCap))
	copy(newSlice, arr.data)
	copy(newSlice[arr.length:], elements)
	arr.data = newSlice

	// set the ADT's length and cap
	arr.updateLenCap()
}
