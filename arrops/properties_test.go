package arrops

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetElement(t *testing.T) {
	arr, _ := CreateArrayADT(0, 5)
	arr.data = append(arr.data, 2, 5, 6, 9, 11)
	value := arr.GetElement(2)
	assert.Equal(t, 6, value)
}

func TestSetElement(t *testing.T) {
	arr, _ := CreateArrayADT(5, 10)
	arr.SetElement(2, 10)
	assert.Equal(t, 10, arr.data[2])
}

func TestIsSorted(t *testing.T) {

	var arr ArrayADT
	var isSorted bool

	// test whether function work for unsorted array
	arr, _ = CreateArrayADT(0, 10)
	arr.data = append(arr.data, 11, 4, 54, 23, 78)
	isSorted = arr.IsSorted(true)
	assert.Equal(t, false, isSorted)

	// test whether function work for sorted ascending array
	arr, _ = CreateArrayADT(0, 10)
	arr.data = append(arr.data, 1, 4, 5, 13, 17)
	isSorted = arr.IsSorted(true)
	assert.Equal(t, true, isSorted)

	// test whether function work for unsorted array
	arr, _ = CreateArrayADT(0, 10)
	arr.data = append(arr.data, 16, 12, 10, 6, 2)
	isSorted = arr.IsSorted(false)
	assert.Equal(t, true, isSorted)
}
