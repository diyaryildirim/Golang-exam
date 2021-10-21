package main

import "fmt"

func main() {
	var i, b2, b3, b4 int
	for i = 1; i <= 100; i++ {
		if i%4==0{
			b4++
		}
		if i%2 == 0 {
			b2++
		} else if i%3 == 0 {
			b3++
		}
	}

	fmt.Printf("2:%d, 3:%d, 4:%d", b2, b3, b4)

}
