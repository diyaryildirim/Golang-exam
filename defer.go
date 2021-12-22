package main

import "fmt"

var isConnected bool = false
func main()  {
	fmt.Printf("Connection open : %v\n", isConnected)
	databaseProcessing()
	fmt.Printf("Cpnnection open : %v\n", isConnected)
	 
}

func databaseProcessing()  {
	connect()
	fmt.Println("Deferring Disconnect!")
	defer disconnect()
	fmt.Printf("Connection Open : %v\n", isConnected)
	fmt.Println("Doing something")
}

func connect()  {
	isConnected = true
	fmt.Println("Connected to Database!")
}

func disconnect()  {
	isConnected = false
	fmt.Println("Disconnected!")
}