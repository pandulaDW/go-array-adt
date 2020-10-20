package arrops

// Max returns the maximum value of the array
func (arr *ArrayADT) Max() int {
	max := arr.data[0]
	for i := 1; i < arr.length; i++ {
		if max < arr.data[i] {
			max = arr.data[i]
		}
	}
	return max
}

// Min returns the minimum value of the array
func (arr *ArrayADT) Min() int {
	min := arr.data[0]
	for i := 1; i < arr.length; i++ {
		if min > arr.data[i] {
			min = arr.data[i]
		}
	}
	return min
}

// Sum returns the sum of the elements of the array
func (arr *ArrayADT) Sum() int {
	sum := 0
	for i := 0; i < arr.length; i++ {
		sum += arr.data[i]
	}
	return sum
}
