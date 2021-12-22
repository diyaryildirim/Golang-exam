package main

import "fmt"

func RecursiveInt(a int) int  {
	if a < 2{
		return a
	}
	return RecursiveInt(a - 1) + RecursiveInt(a - 2)
}

func main()  {
	fmt.Println(RecursiveInt(6))
}
