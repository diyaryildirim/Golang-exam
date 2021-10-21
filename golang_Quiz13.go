package main

import (
	"fmt"
)

func main() {
	var y int
	for y, z := 1, 1; y < 10; y++ {
		_ = y
		_ = z
	}
	fmt.Println(y)
}
