package main

import (
	"fmt"
	"math/rand"

	"github.com/pandulaDW/go-array-adt/arrops"
)

func main() {
	a, _ := arrops.CreateArrayADT(0, 10)

	a.Push(2, 4, 5, 13, 15, 9, 10, 11, 18)
	a.PrintData()

	a.RotateRight()
	a.RotateRight()
	a.RotateRight()
	a.PrintData()

	for i := 0; i < 10; i++ {
		rand.Seed(int64(i))
		fmt.Print(rand.Intn(100), ", ")
	}
}
