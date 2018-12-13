// 1.
-	var x string = "Hello E-phraim"

-	x := "Hello E-phraim"

// 2.
package main

import "fmt"

func main() {
	x := 5; x +=1
	//this implies the value of x which is 5 + 1 ..the answer will be 6
}

// 3.
// Scope is the range of places where you allowed to use a variable.
// A scope of a variable can come in two ways ..First if you define the variable
// outside the main function it'll have a large or limitless scope, it can called from anywhere
// Second if a variable is defined in a function it's scope is within the curly braces.

// the var way of creating variables is that it can be changed or manipulated later within
// the code while the const way of creating is the opposite, you can't change or
// manipulate it later.

// 5.
package main

import "fmt"
func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5/9

	fmt.Println(output)
}

// 6.
package main

import "fmt"
func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input*0.3048)
	fmt.Println(output)
}


