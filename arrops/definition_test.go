package arrops

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateADT(t *testing.T) {
	length, capacity := 5, 10
	arr, _ := CreateArrayADT(length, capacity)

	// assert length
	assert.Equal(t, length, arr.length)

	// assert capacity
	assert.Equal(t, capacity, arr.capacity)

	// assert elements
	assert.ElementsMatch(t, make([]int, length, capacity), arr.data)
}

func TestGetLength(t *testing.T) {
	arr, _ := CreateArrayADT(5, 10)
	assert.Equal(t, 5, arr.GetLength())
}

func TestGetCapacity(t *testing.T) {
	arr, _ := CreateArrayADT(5, 10)
	assert.Equal(t, 10, arr.GetSize())
}

func TestUpdateLenCap(t *testing.T) {
	arr, _ := CreateArrayADT(5, 10)
	arr.data = make([]int, 10, 15)
	arr.updateLenCap()

	// assert if length and capacity has been updated properly after a direct assignment
	assert.Equal(t, 10, arr.length)
	assert.Equal(t, 15, arr.capacity)
}
