package main

import "fmt"

type human struct {
	FirstName string
	LastName string
	Age int
}

func main()  {
	a := human{FirstName: "Diyar"}
	fmt.Println(a.FirstName)
}