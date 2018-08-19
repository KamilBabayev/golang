package main

import (
	"fmt"
)

func main(){
	user := User{name:"Kamil", sname:"Babayev"}
	fmt.Println(user.name, user.sname)
    fmt.Println(createUserFullName(user.name, user.sname))
	fmt.Println(calc(3,44))
	fmt.Println(my_func(2.3, 33.4))
}

type User struct{
	name string
    sname string
	city string
}

func createUserFullName(name string, sname string) string {
	return name + " and " + sname
}

func calc(a int, b int) int {
	return a + b
}

func my_func(a float32, b float32) float32 {
	return a * b
}
