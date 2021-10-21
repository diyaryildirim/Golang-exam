package main

import "fmt"

func main()  {
	test := []string{"a" , "b" , "c"}
	test = append(test[1:], test[:1]...)
	fmt.Println(test)
}
