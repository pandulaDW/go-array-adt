package arrops

import (
	"errors"
	"fmt"
)

// ArrayADT is an ADT built around go builtin arrays
type ArrayADT struct {
	data     []int
	length   int
	capacity int
}

// CreateArrayADT initializes the array ADT with the given len and cap
func CreateArrayADT(len, cap int) (ArrayADT, error) {
	if len > cap {
		return ArrayADT{}, errors.New("length cannot be greater than the capacity")
	}
	newArr := ArrayADT{make([]int, len, cap), len, cap}
	return newArr, nil
}

// PrintData prints the data, length and capacity of the underlying array
func (arr *ArrayADT) PrintData() {
	fmt.Printf("%v -> len: %d, cap: %d\n", arr.data, arr.length, arr.capacity)
}

// update the length and cap of ADT using the current underlying array
func (arr *ArrayADT) updateLenCap() {
	arr.length = len(arr.data)
	arr.capacity = cap(arr.data)
}

// GetLength returns the length of the array
func (arr *ArrayADT) GetLength() int {
	return arr.length
}

// GetSize returns the capacity of the array
func (arr *ArrayADT) GetSize() int {
	return arr.capacity
}
