package main

import "fmt"

func main() {
	// :show start
	// Creating arrays of 6 elements of type int,
	// and put elements 1, 2, 3, 4, 5 and 6 inside it, in this exact order:
	var array1 = [6]int{1, 2, 3, 4, 5, 6}   // classical way
	var array2 = [6]int{1, 2, 3, 4, 5, 6}   // a less verbose way
	var array3 = [...]int{1, 2, 3, 4, 5, 6} // the compiler will count the array elements by itself

	fmt.Println("array1:", array1) // > [1 2 3 4 5 6]
	fmt.Println("array2:", array2) // > [1 2 3 4 5 6]
	fmt.Println("array3:", array3) // > [1 2 3 4 5 6]

	// Creating arrays with default values inside:
	zeros := [8]int{}       // Create a list of 8 int filled with 0
	ptrs := [8]*int{}       // a list of int pointers, filled with 8 nil references ( <nil> )
	emptystr := [8]string{} // a list of string filled with 8 times ""

	fmt.Println("zeroes:", zeros)      // > [0 0 0 0 0 0 0 0]
	fmt.Println("ptrs:", ptrs)         // > [<nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>]
	fmt.Println("emptystr:", emptystr) // > [       ]
	// values are empty strings, separated by spaces,
	// so we can just see separating spaces

	// Arrays are also working with a personalized type
	type Data struct {
		Number int
		Text   string
	}

	// Creating an array with 8 'Data' elements
	// All the 8 elements will be like {0, ""} (Number = 0, Text = "")
	structs := [8]Data{}

	fmt.Println("structs:", structs) // > [{0 } {0 } {0 } {0 } {0 } {0 } {0 } {0 }]
	// prints {0 } because Number are 0 and Text are empty; separated by a space
	// :show end
}
