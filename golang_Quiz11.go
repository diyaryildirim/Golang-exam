package main

import "fmt"

func main()  {
	namestring := "dev"
	for _, ch := range []rune (namestring){
		fmt.Println(ch-'a')
	}
}
