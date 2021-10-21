package main

import "fmt"

func main()  {
	var letters =[]string{"a", "b" , "c"}
	var someLetters= make([]string,2)
	copy(someLetters, letters)
	fmt.Println(someLetters)
}
