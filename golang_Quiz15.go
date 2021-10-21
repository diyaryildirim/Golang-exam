package main

import "fmt"

type data interface {
	input()
	output()
}
type inptoutputhandler struct {}

var (
	ch = make(chan string, 1)
)
func (io inptoutputhandler) input(){
	ch <- "hello"
}
func (io inptoutputhandler) output(){
	fmt.Println(<-ch)
}
func operation(d data) {
	d.input()
	d.output()
}
func main()  {
	var dataCaller inptoutputhandler
	operation(dataCaller)
}