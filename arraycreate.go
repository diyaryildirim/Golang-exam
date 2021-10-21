package main

import "fmt"

func main()  {
	var cars [3]string
	cars[0] ="Tesla"
	cars[1] ="BMW"
	cars[2] ="Mercedes"

	for _, value := range cars {
		fmt.Println( value)
	}
}