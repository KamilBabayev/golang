package main

import (
	"fmt"
)

func main(){
	d1 := "kamil"
	d2 := "famil"
	d3 := "yunis"
	fmt.Println("Hello World!!!")
	printData(d1, d2, d3)
}

func printData(data ...string) {
	//for _,  name := range data {
	for index,  name := range data {
		fmt.Println(index, name)
	}
}
