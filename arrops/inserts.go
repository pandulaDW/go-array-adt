package arrops

import (
	"math"
)

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

// InsertSorted insert elements by preverserving the sort order. Acts as a set and
// only unique elements will be inserted
//
// It assumes that the elements are sorted by ascending order and that there's enough capacity
func (arr *ArrayADT) InsertSorted(value int) {

	// if value is higher than maximum, just append the value to the end
	if arr.data[arr.length-1] < value {
		arr.Push(value)
		arr.updateLenCap()
		return
	}

	// if value is lower than the minimum, shifts all other elements to right and insert at the beginning
	if arr.data[0] > value {
		arr.data = arr.data[:arr.length+1]
		copy(arr.data[1:], arr.data[0:])
		arr.data[0] = value
		arr.updateLenCap()
		return
	}

	low, high := 0, arr.length
	mid := (low + high) / 2

	for low <= high {
		if arr.data[mid] == value {
			return
		}
		if arr.data[mid] > value {
			high = mid - 1
		}
		if arr.data[mid] < value {
			low = mid + 1
		}
		mid = (low + high) / 2
	}

	// switch the low and high values for clarity
	low, high = high, low

	// extend the slice
	arr.data = arr.data[:arr.length+1]

	// copy the original elements from the high position to the end of the slice
	copy(arr.data[high+1:arr.length+1], arr.data[high:])

	// set the high position element the inserted value
	arr.data[high] = value

	// update the length and capacity
	arr.updateLenCap()
}

// MergeArray merges another array/slice with the current array and
// maintains the array order. Can be considered as a set operation
func (arr *ArrayADT) MergeArray(otherArr []int) {

	var a, b []int
	c := make([]int, arr.length+len(otherArr))

	// assign aliases to long and short array
	if arr.length >= len(otherArr) {
		a, b = arr.data, otherArr
	} else {
		a, b = otherArr, arr.data
	}

	for i, j, k := 0, 0, 0; i < len(c); k++ {
		if j == len(b) {
			copy(c[k:], a[i:])
			break
		}

		if a[i] < b[j] {
			c[k] = a[i]
			i++
			continue
		}
		if a[i] > b[j] {
			c[k] = b[j]
			j++
			continue
		}
	}

	arr.data = c
	arr.updateLenCap()
}
