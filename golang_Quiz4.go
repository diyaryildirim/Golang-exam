package main

import (
	"fmt"
	"time"
)
func printDetail() {
	for i := 0; i <= 230; i++ {
		fmt.Println(i)
		time.Sleep(10 * time.Millisecond)
	}
}
func printSupport() {
	printDetail()
}
func main() {
	printSupport()
	time.Sleep(2 * time.Second)
	fmt.Println( "Hello, playground")
}