package arrops

import "testing"

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}

// TestCreateADT tests whether ADT is created correctly
func TestCreateADT(t *testing.T) {
	length, capacity := 5, 10
	arr, _ := CreateArrayADT(length, capacity)
	if arr.length != length {
		t.Fail()
	}
	if arr.capacity != capacity {
		t.Fail()
	}
	if !compareSlices(arr.data, make([]int, 5)) {
		t.Fail()
	}
}

// TestCreateADT tests whether the length is returned correctly
func TestGetLength(t *testing.T) {
	arr, _ := CreateArrayADT(5, 10)
	if arr.GetLength() != 5 {
		t.Fail()
	}
}
