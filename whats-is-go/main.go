package main

import (
	"fmt"
)

//User ...
type User struct {
	ID int
	Name string
	Phone string
}

//Developer ...
type Developer struct{
	user User 
	Langs []string
}

func main() {
	var num int = 4
	var num2 int = 5
	fmt.Println(num + num2)

	//------------------------------

	var num3, num4 int = 5, 6
	fmt.Println(num3 + num4)

	//------------------------------

	var num5, num6 = 7, 8
	fmt.Println(num5 + num6)

	//-------------------------------

	num7, num8 := 9, 10
	fmt.Println(num7 + num8)

	//------------------------------

	user := User{1, "majid", "0916970611"}
	fmt.Println(user)
	fmt.Println(user.ID, user.Name, user.Phone)

	//And we can use Array, Slices, Maps 
	//to structure different formats and etc

	developer := Developer{user, []string{"Go", "Rust", "Elixire", "Python"}}

	fmt.Println(developer)
}