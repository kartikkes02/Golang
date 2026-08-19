// Generics -> a way to define a function or a type that can work with different types.
// Generics are useful for creating reusable code that can work with different types without
// sacrificing type safety.
// In Go, we can use the type parameters to define generics.
// Type parameters are defined using square brackets [] and can be used to define a function or a type that can work with different types.
// Type parameters can be used to define a function or a type that can work with different types.
// It's under in the 1.18 version of Go.

package main

import "fmt"

func printSlices[T any](items []T) {
	// if ony pass the string and int so,
	// func printSlices[T string | int](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {

	// why we use generics?
	// not again and again write the same code for different types.
	fmt.Println("Generics")

	names := []string{"apple", "banana", "cherry"}
	printSlices(names)

}
