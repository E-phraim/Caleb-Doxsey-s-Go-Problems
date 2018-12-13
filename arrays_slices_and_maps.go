// 1.
// package main

// import "fmt"

// func main () {
// 	x := [5]float64{100, 200, 300, 400, 500}
// 	fmt.Println(x[4])
// }

// In the above code we have 5 elements, if we run this code we'll have "500" as our
// answer because it is the 4th indexed element in that array.

// 2.

// The size of the slice is 9. For further knowlege it'll be nice to know that the 
// make function allows for a 3rd parameter. Now let's look at this way, we have the array
// type as 'float64' the next parameter can be the minimum capacity and then 9 the
// higher value taken as the maximum capacity.

// 3.
// package main

// import "fmt"

// func main () {
// 	x := [6]string{"a", "b", "c", "d", "e", "f"}
// 	fmt.Println(x[2:5])
// }

// the logic behind these ':' expressions are quite simple...it prints out everything
// from the 2nd element(start counting from 0) till the fourth element (now that is the nature of this
// expression...it prints everything and stops at the last element before the 5th element).
// Our answer will be [c d e] 2nd, 3rd, 4th, !5th.

// 4.
//This problem is a bit complex, please follow step-by-step
package main 
import "fmt"

func main () {
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	smallest := x[0] //set the smallest number to the first element of the list
	for _, num := range x[1:] { //iterate(loop) over the rest of the list
		if num < smallest { //if num is smaller than the current smallest number
			smallest = num //set the smallest to num
		}
	}
	fmt.Println(smallest)
}