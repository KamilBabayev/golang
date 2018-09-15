package main

import (
	"fmt"
//	"os"
)



func main () {
	c := Circle{x: 1.1, y: 2.2, r: 1.1}
	fmt.Println(c.x,
				c.y, 
				c.r,
	)
    a := Car{a: 100, b:200, c:300}
    w := [] int{a.a, a.b, a.c}
    for _, i := range w {
		fmt.Println(i)
	}
	
}


type Circle struct  {
    x float64
	y float64
	r float64
}

type Car struct {
	a, b, c int
}
