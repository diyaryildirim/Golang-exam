package main

import "fmt"

func main()  {
	slice:= []int{1,56,27,98,0}

	for i, value :=range slice{
		fmt.Printf("index %d Değer : %d\n", i , value)
	}

	var numbers = make([]int, 5, 5)
	numbers[0] = 345
	numbers[1] = 23
	numbers[2] = 35
	numbers[3] = 45
	numbers[4] = 546
	fmt.Println(numbers)
	numbers = append(numbers, 789)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))
}
