package main

import (
	"fmt"
	"strconv"
)

func main()  {

	// User data create

	fmt.Println("User create V1")

	user1 :=&User{
		ID:   1,
		FirstName:   "Diyar",
		LastName:     "Yildirim",
		UserName:    "Diyaryildirim",
		Age: 21,
		Pay: &Payment{
			Salary: 3.555,
			Bonus: 650,
		},
	}

	// fmt.Println(user1)
	fmt.Println(user1.GetFullName())
	fmt.Println(user1.GetPayment())
	fmt.Println("Maaş : " + strconv.FormatFloat(user1.GetPayment(), 'g', -1 , 64))



	//User Create V2
	fmt.Println("User Create V2")

	user2 := NewUser()
	user2.FirstName = "Bilal"
	user2.LastName = "Ekmis"
	user2.Age = 20
	user2.UserName = "Casper"

	user2.Pay = &Payment{Salary: 5.550, Bonus: 850,}
	fmt.Println(user2.GetFullName())
	fmt.Println(user2.GetPayment())
	fmt.Println("Maaş : " + strconv.FormatFloat(user2.GetPayment(), 'g', -1, 64))



}

type User struct {
	ID  int
	FirstName string
	LastName string
	UserName string
	Age     int
	Pay    *Payment
}
func NewUser() *User{
	u := new(User)
	u.Pay = NewPayment()
	return u
}
type Payment struct {
	Salary float64
	Bonus float64
}
func NewPayment() *Payment{
	p := new(Payment)
	return p
}

func (u User) GetFullName() string  {
	return u.FirstName+ " " + u.LastName
}

func(u *User) GetUserName() string{
	return u.UserName
}

func (u *User) GetPayment() float64{
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
}