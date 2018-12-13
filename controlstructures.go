// 1.
package main

import "fmt"
func main() {
	i := 10
	if i > 10 {
		fmt.Println("Big")
	}	else {
		fmt.Println("Small")
	}
}

// Your answer should be "Small" because i(10) is not bigger than 10

// 2.

package main
import "fmt"

func main () {
	for i := 1; i <= 100; i++{
		if i % 3 == 0{
			fmt.Println(i)
		}
	}
}

// 3.
package main
import "fmt"

func main () {
	for i := 1; i <= 100; i++{
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
