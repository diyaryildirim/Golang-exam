package main

import "fmt"

func main()  {
	aday := [6]int {
		81,
		100,
		27,
		95,
		45,
		68,
	}
	var adaysayisi int = len(aday)
	for i:=0; i < adaysayisi; i++{
		if aday[i] <50 {
			fmt.Println(i,". aday başarısız oldu :" , aday[i])
		}
	}
}
