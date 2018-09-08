package main

import (
	"fmt"
)


func main() {
	names := make(map[string]int)
	names["kamil"] = 2323
	names["famil"] = 1212
	names["emin"] = 111
	fmt.Println(names)
	fmt.Println(names["famil"])

    if name, ok := names["kamil"]; ok {
	  fmt.Println(name, ok, "Check if kamil exist in this map")
	}

    delete(names, "kamil")
	fmt.Println(names)

    if name, ok := names["kamil"]; ok {
	  fmt.Println(name, ok)
	}

    test := map[string]string{
		"user1":"Admin",
		"user2":"operator",
		"user3":"system",
	}
    fmt.Println(test)
	second()
   
}

func second() {
    fmt.Println()
	fmt.Print("This is second function inside code beside main func\n")
}

