package main

import (
	"fmt"

	"C" //cheat antivirus

	"my.com/calc"
)

func main() {

	var i = 2.233
	var j = 3.712
	var result = calc.Sub(i, j)
	fmt.Printf("Subtracting %f and %f = %f ", i, j, result)
}
