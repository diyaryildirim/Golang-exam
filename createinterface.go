package main

import (
	"fmt"
	"strconv"
)

func CarExecute(c Carface)  {
	fmt.Println("\n")
	fmt.Println("Araç Bilgi : \n" +c.Information())
	fmt.Println("\n")

	msg:= " "

	isRun := c.Run()
	if isRun{
		msg = "Çalışıyor"
	}else {
		msg = "Çalışmıyor"
	}
	fmt.Println("Araç" + msg + ".")

	isStop := c.Stop()
	if isStop{
		msg = "Durdu"
	}else {
		msg = "Duramadı. Fren Tutmuyorr!!!"
	}
	fmt.Println("Araç" + msg + ".")

}

func main()  {

	//FERRARİ
	ferr := new(Ferrari)
	ferr.Brand = "Ferrari"
	ferr.Model = "599 GTo"
	ferr.Color = "Ferrari's RED"
	ferr.Speed = 335
	ferr.Price = 3.4
	ferr.Special= true
	//fmt.Println(ferr.information())


	//Mercedes
	merc := new(Mercedes)
	merc.Brand ="Mercedes"
	merc.Model = "G"
	merc.Color = "Metallic Gray"
	merc.Speed = 275
	merc.Price = 1.5
	//fmt.Println(merc.information())

	CarExecute(ferr)
	CarExecute(merc)

}
//interface

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

// Base Structs

type Car struct {
	Brand   string
	Model  string
	Color  string
	Speed  int
	Price  float64
}

type SpecialProduction struct {
	Special  bool
}

// Ferrari

type Ferrari struct {
	Car  // composition
	SpecialProduction
}

func (_  Ferrari) Run() bool{
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (x Ferrari) information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color : "+ x.Color + "\n" + "\t" + "Speed : " + strconv.Itoa(x.Speed)+ "\n" + "\t" + "Price : " + strconv.FormatFloat(x.Price, 'g',-1, 64)+ "Million"
	add := "Yes"
	if x.Special{
		ret += "\n" + "\t" + "Special : " + add
	}
	return  ret
}

//Lamborghini

type Lamborghini struct {
	Car  // composition
	SpecialProduction
}

func (_  Lamborghini) Run() bool{
	return true
}

func (_ Lamborghini) Stop() bool {
	return true
}

func (x Lamborghini) information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color : "+ x.Color + "\n" + "\t" + "Speed : " + strconv.Itoa(x.Speed)+ "\n" + "\t" + "Price : " + strconv.FormatFloat(x.Price, 'g',-1, 64)+ "Million : "
	add := "Yes"
	if x.Special{
		ret += "\n" + "\t" + "Special : " + add
	}
	return  ret
}

//MERCEDES

type Mercedes struct {
	Car  // composition
	SpecialProduction
}

func (_  Mercedes) Run() bool{
	return true
}

func (_ Mercedes) Stop() bool {
	return true
}

func (x Mercedes) information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color : "+ x.Color + "\n" + "\t" + "Speed : " + strconv.Itoa(x.Speed)+ "\n" + "\t" + "Price : " + strconv.FormatFloat(x.Price, 'g',-1, 64)+ "Million : "
	return  ret
}


