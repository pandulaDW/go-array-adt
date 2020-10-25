package main

import (
	"github.com/pandulaDW/go-array-adt/arrops"
)

func main() {
	a, _ := arrops.CreateArrayADT(0, 10)

	a.Push(1, 4, 5, 9, 12, 15)

	a.InsertSorted(10)
	a.InsertSorted(3)
	a.InsertSorted(20)

	a.PrintData()
}
