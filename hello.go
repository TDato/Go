package main

// Declare a main package
// a package is a way to group functions
// it's made up of all the files in the same directory
/**
 * import the popular 'fmt' package
 * fmt contains functions for formatting the text,
 * including printing to the console.
 *
 * The fmt package is one of the standard library packages
 * that is installed when you install Go
 */
import (
	"fmt"
	"time"

	"math/rand"

	"rsc.io/quote"
)

/**
* Implement a main function to print a message to the console.
* A main function executes by default when you run the main package.
 */
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(quote.Go())
	fmt.Println("My favorite number is", rand.Intn(19))
	fmt.Println(add(555, rand.Intn(111)))
	var x int = rand.Intn(555)
	var y int = rand.Intn(111)
	x, y = split(add(x, y))
	fmt.Printf("after split:\nx: %v, y: %v\n", x, y)
	fmt.Println()
}

//func add(x int, y int) int
// when two or more consecutive named function parameters
// share a type, you can omit the typ from all but the last.
func add(x, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
