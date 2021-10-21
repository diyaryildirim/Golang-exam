package main

import (
	"fmt"
	"strings"

	color "github.com/fatih/color"
)

func main()  {

	fmt.Println(strings.HasPrefix("unity_test", "e"))

	color.red("error messgae")
}