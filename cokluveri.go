package main

import "fmt"

func main()  {
	nokta := [3][2]int{{11,12}, {21,22}, {31,32}}
	fmt.Println("Dizinin Tamamı :", nokta)
	fmt.Println("\n3x2 Matris biçiminde:")
	for satir :=0; satir<3; satir++{
			for sutun :=0; sutun<2 ; sutun ++{
				fmt.Print(nokta[satir][sutun], " ")
			}
			fmt.Println()
	}
}
