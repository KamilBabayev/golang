package main	//Package declaration, libraries
// Comment line, lines starts or contains with // is comments
/*
This is multiline comment form,
Go also supports this type, 
but, I miss for python :)
*/
import "fmt"	//we import fmt package (formatting for input, output) files in fmt package starts with fmt.

/*
Function declaration, every function starts with func keywork, func name and arguments
main function is special and it is called when program executes, that is why we see print result.
*/
func main() {		
	fmt.Println("Hello World!! ")   // Println  takes one argument and executes it.
}

func new() {
	fmt.Println("This line will not print, because it is just declaration")
}





// go run first.go		- this is just for running script
// go build first.go		- this will compile script and create binary file
// go doc frm  or go doc frm Println	 -  this will print help menu about frt and its internal method Println
