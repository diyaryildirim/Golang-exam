package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	max = 5
	min = 0
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	for i := 0; i < 2; i++ {
		fmt.Println(rand.Intn(max-min) + min)
	}
}
