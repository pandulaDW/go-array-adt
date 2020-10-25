package main

import (
	"fmt"

	"github.com/pandulaDW/go-array-adt/arrops"
)

func main() {
	a, _ := arrops.CreateArrayADT(0, 10)

	a.Push(11, 4, 54, 23, 78)

	fmt.Println(a.IsSorted(true))
}
