package main

import (
	"fmt"

	"./arrops"
)

func main() {
	a, _ := arrops.CreateArrayADT(0, 10)

	a.Push(2, 4, 5, 7, 9, 13, 17, 19, 25)
	a.PrintData()

	fmt.Println(a.BinarySearch(1))
}
