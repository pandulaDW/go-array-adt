package main

import (
	"github.com/pandulaDW/go-array-adt/arrops"
)

func main() {
	a, _ := arrops.CreateArrayADT(0, 4)
	a.Push(3, 8, 16, 20, 25)

	otherArr := []int{4, 10, 12, 22, 23}

	a.MergeArray(otherArr)

	a.PrintData()
}
