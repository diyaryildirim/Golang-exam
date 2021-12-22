package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Su anki unix zaman : %v\n ", time.Now().Unix())
	time.Sleep(20 * time.Second)
	fmt.Println("Suanki Unix Zaman : %v\n ", time.Now().Unix())
}
