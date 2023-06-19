package main

import (
	"fmt"

	"C"

	"my.com/calc"
)

func main() {

	var (
		i      = 2
		j      = 3
		result = calc.Add(i, j)
	)
	fmt.Printf("Adding %d and %d = %d ", i, j, result)
}
