package arrops

// Reverse the array in-place
func (arr *ArrayADT) Reverse() {
	reversedArr := make([]int, arr.length, arr.capacity)
	for i := 1; i <= arr.length; i++ {
		reversedArr[i-1] = arr.data[arr.length-i]
	}
	arr.data = reversedArr
}

// ShiftLeft shifts the array by one element to the left
func (arr *ArrayADT) ShiftLeft() {
	for i := 0; i < arr.length-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.data[arr.length-1] = 0
}

// ShiftRight shifts the array by one element to the right
func (arr *ArrayADT) ShiftRight() {
	shiftedArray := make([]int, arr.length, arr.capacity)
	shiftedArray[0] = 0
	for i := 1; i < arr.length; i++ {
		shiftedArray[i] = arr.data[i-1]
	}
	arr.data = shiftedArray
}

// RotateLeft rotates the elements to the left by a single shift
func (arr *ArrayADT) RotateLeft() {
	firstEl := arr.data[0]
	arr.ShiftLeft()
	arr.data[arr.length-1] = firstEl
}

// RotateRight rotates the elements to the right by a single shift
func (arr *ArrayADT) RotateRight() {
	lastEl := arr.data[arr.length-1]
	arr.ShiftRight()
	arr.data[0] = lastEl
}

// ShiftsNegative shifts negative elements to the left
func (arr *ArrayADT) ShiftsNegative() {
	for i, j := 0, arr.length-1; i < arr.length; {

		// loop until a positive element is found from the end side
		if arr.data[j] > 0 {
			j--
		}

		// loop until a negative element is found from the starting side
		if arr.data[i] < 0 {
			i++
		}

		// terminating condition
		if i > j {
			return
		}

		// if a negative element is found from the starting side, swap it with the end side
		// if the end is still isn't positive, keep iterating
		if arr.data[i] > 0 {
			if arr.data[j] > 0 {
				continue
			}
			arr.data[i], arr.data[j] = arr.data[j], arr.data[i]
			i++
			j--
		}
	}
}
