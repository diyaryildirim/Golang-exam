package main

import (
	"fmt"
)

func main() {

	var x int
	lstArray := [3]int{1, 2, 3}
	x, lstArray = lstArray[0], lstArray[1:]
	fmt.Println(
		x,
		lstArray,
	)
}
