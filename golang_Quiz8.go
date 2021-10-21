package main

import(
	"fmt"
)

func printChannel(ch chan int){
	for{
		select{
		case num := <-ch:
			fmt.Printf("%d " , num)
		}
	}
}

func main(){
	ch := make (chan int)
	for i := 0; i<5; i++{
		ch <- i
	}
	go printChannel(ch)
}