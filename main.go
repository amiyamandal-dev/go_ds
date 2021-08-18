package main

import (
	"fmt"

	"work_code/movenegetivenumber"
)

func main() {
	fmt.Println("hello world")
	A := [...]int{-1, 7, -11, 10, -111, 4, -3434, 3, 20, 15, -9999}
	movenegetivenumber.MoveNegetiveNumber(A[:])
	fmt.Println(A)
}
