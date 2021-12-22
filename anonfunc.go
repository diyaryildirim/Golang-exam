package main

func main()  {
	 addFunc := func(terms ...int)(numTerms int, sum int){
		 for _, term:=range terms{
			 sum += term
		 }
		 numTerms = len(terms)
		 return
	 }
	 numTerms, sum :=addFunc(2,5)
	 println("Added :", numTerms, "terms fora total of", sum)
}