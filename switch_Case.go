package main

import "fmt"

func main() {
	sayi := 0
	switch {
	case sayi > 0:
		fmt.Println("pozitif")
	case sayi == 0:
		fmt.Println("nötr")
	default:
		fmt.Println("negatif")
	}
}
