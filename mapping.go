package main

import "fmt"

func main()  {

	// 1.
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[43] = "Poo"
	myMap[13] =  "Repair"
	fmt.Println(myMap)

	//2.
	states := make (map[string]string)
	states["IST"] = "Istanbul"
	states["ANK"] = "Ankara"
	states["ANT"] = "Antalya"
	fmt.Println(states)

	delete(states, "ANT")
	fmt.Println(states)

	states["ERZ"] = "Erzurum"

	keys := make([]string, len(states))
	i := 0
	for k := range states{
		keys[i] = k
		i++
	}
	fmt.Println(keys)

}
