package main

import (
	"fmt"

	"github.com/pandulaDW/go-array-adt/arrops"
)

func main() {
	a, _ := arrops.CreateArrayADT(0, 10)

	a.Push(10, 9, 5, 3)

	fmt.Println(a.IsSorted(false))
}
