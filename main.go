package main

import (
	"fmt"

	"./arrops"
)

func main() {
	a, _ := arrops.CreateArrayADT(0, 10)

	a.Push(1, 2, 3, 4, 5, 6, 7, 8, 9)
	a.PrintData()

	fmt.Println(a.LinearSearch(5))  // 2
	fmt.Println(a.LinearSearch(99)) // -1
	fmt.Println(a.Contains(91))

	a.Delete(2)
	a.PrintData()
}
