package main

import "fmt"
// Go is statically typed lang, variables have types and that type can not change.
// GO has built-in data types. integers and float-poit numbers
// Integer types are uint8,uint16,uint32,uint64 and int8,int16,int32,int64 (2 alias types : byte(8bit) - uint8, rune - int32)
// uint is unsigned integers which conrains positive numbers or zero. floar point is 1.23, 33.5, 0.23.

func main() {
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1 + 1 =", 1.0 + 1.0)	
	fmt.Println("1 + 1 =", 1.3 + 1.2)	
	fmt.Println(len("hello world"))
	fmt.Println("hello world"[1])
	fmt.Println("hello" + "world")
}

// Go lang supports this standard arithmetic operations: +, -, *, /, %
// Golang support string type. string is sequence of characters
// Strings are made of individual bytes, usually one for each character
// String literals can be created via "string" or `string` backticks
// len("hello")   - finds lenght,  "hello"[1]  - second character
// "hello" + "world"  - concat
