package main

import (
	"fmt"
)

type Option interface {
	firstCall(*funcOption)
}

type funcOption struct {}

func (opt *funcOption) firstCall(do *funcOption)  {
	fmt.Println("First call")
}

func main()  {
	var opt Option
	data := &funcOption{}
	opt = data
	opt.firstCall(data)
}