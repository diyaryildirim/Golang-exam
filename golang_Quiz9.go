package main

import(
	"fmt"
	"time"
)

func printNumber(done <-chan int) {
	for i:=0; i<10; i++{
		select{
		case <- done :
			fmt.Println("Time to close")
			return
		default:
			fmt.Println(i)
			time.Sleep(1*time.Second)
		}
	}
}

func main(){
	done := make(chan int , 2)
	defer close(done)
	go func (){
		printNumber(done)
	}()
	time.Sleep(3 * time.Second)
	done <- 1
	time.Sleep(3* time.Millisecond)
	fmt.Println("Exited")
}