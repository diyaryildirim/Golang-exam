package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2021, time.December, 10, 15, 0, 0, 0, time.UTC)
	fmt.Println("Çalışıyor : %s\n", t)

	fmt.Println("***********************")

	now := time.Now()
	fmt.Println("Mevcut Zamanı: %s\n", now)

	fmt.Println("***********************")

	fmt.Println("Ay :", now.Month())
	fmt.Println("Gün : ", now.Day())
	fmt.Println("Haftanın Günü : ", now.Weekday())

	fmt.Println("***********************")

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Println("Tomorrow is %v %v %v %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	fmt.Println("***********************")
	longFormat := "Monday, December 20 , 2021"
	fmt.Println("Tomorrow is ", tomorrow.Format(longFormat))

}
