package main // Declares the main package. A Go program starts from the main function in this package.

import "fmt" // Imports the "fmt" package from the standard library. Go imports entire packages, not individual functions.

func main() {
	fmt.Printf("Hello %s!\n", "World")
}
