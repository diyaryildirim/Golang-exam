package main

import (
	"fmt"
	"strings"
)
func main(){
	var str strings.Builder

	str.WriteString("a")
	str.WriteString("b")
	str.WriteString("c")

	fmt.Println(str.String())
}