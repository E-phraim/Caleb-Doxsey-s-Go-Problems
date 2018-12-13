// 1. Positive integers are generally stored as simple binary numbers
//  (1 is 1, 10 is 2, 11 is 3 and so on). Negative integers are stored 
//  as the two's complement of their absolute value. 


 //3.
package main

import "fmt"

func main() {
	fmt.Println(321325 * 424521)
}

//4.
// A string is a sequence of characters with a definite length used to represent text.
// Go strings are made up of individual bytes, usually one for each character.

// Finding the length of a string can done by len("Hello World")

// example:

package main

import "fmt"

func main() {
	fmt.Println(len("Hello World"))
}

//5.
// Question is  What's the value of the expression 
// (true && false) || (false && true) || !(false && false)?

true && false = false
false && true = false
!(false && false) = true

false || false || true
true || true = true 