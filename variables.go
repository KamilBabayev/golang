package main

import (
	"fmt"
    "strconv"
)

var User1 string = "Kamil"
var User2 = "Yunis"
var User4, User5 = "Admin", "Demo"
const var10 = "This is demo"
const name, sname = "Kamil", "Babayev"
var age int = 111

var num = 1


func main() {
    var mystring = "1111111111"
    User6 := "demo"
    desc := "This is just description of the more important matter"
	fmt.Println(User1)
	fmt.Println(User2)
    fmt.Println(User1, User2)
    fmt.Println(User4, User5, User6)
    fmt.Println(var10)
    fmt.Println(name, sname)
    fmt.Println(age, "    ", desc)
    fmt.Println(num)
    number, err := strconv.Atoi(mystring)
    fmt.Println(number, err)
    fmt.Printf("%v \n", number)
    fmt.Printf("%T \n", number)
    fmt.Printf("%T \n", desc)
}

