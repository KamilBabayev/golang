package main

import (
	"fmt"
)

func main() {
	fmt.Println("Helow World!")
    user := User{name:"Kamil", sname: "Babayev", city: "Baku", position: "DevOps"}
    pos := Position{name: "linux engineer",  requirements: "scrpting, coding, ops", team: "prod team"}
    fmt.Println(user.name, user.sname, user.city, user.position)
    fmt.Println(pos.name, pos.requirements, pos.team)
    pos.name = "Frontend Developer"
    fmt.Println(pos.name)
}

type User struct {
	name string
    sname string
	city string
    position string
}

type Position struct {
	name string
	requirements string
	team string
}
