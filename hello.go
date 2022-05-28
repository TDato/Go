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
	fmt.Println(add(555, 111))
}

func add(x int, y int) int {
	return x + y
}
