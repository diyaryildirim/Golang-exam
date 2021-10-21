package main

import "fmt"

func main()  {
	baskentler := map[string]string{"TURKEY": "Ankara", "Japan": "Tokyo", "Chinese": "Hong kong", "U.S.A" : "Washington D.C", "Ukraine ": "Kyiev"}
	for key, val:= range baskentler{
		fmt.Println("Map item: Capital of ", key, "is",val)
	}
}