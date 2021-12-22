package main

func main()  {
	numTerms, sum := add(3,4,5)
	println("Added : ",numTerms, "\nterms for a total of", sum )
}

func add(terms ...int) (int,int)  {
	result :=0
	for _,term :=range terms{
		result += term
	}
	return len(terms), result
}