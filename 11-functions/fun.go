// // package main

// // import "fmt"

// // func main() {
// // 	output("Hello World")
// // }

// // func output(a string) {
// // 	fmt.Print(a)
// // }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// output(2.2, 5.0, 10.0)
// 	fmt.Println("Functions")
// }

// // func output(investmentAmount float64, expectedReturnRate float64, years float64) (futureValue float64) {
// // 	fmt.Println("Investment calculator: ")
// // 	fmt.Print("Enter Investment Amount: ")
// // 	fmt.Scan(&investmentAmount)

// // 	fmt.Print("Enter Expected Return Rate: ")
// // 	fmt.Scan(&expectedReturnRate)

// // 	fmt.Print("Enter Number of Years: ")
// // 	fmt.Scan(&years)

// // 	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// // 	return futureValue
// // }

package main

import "fmt"

func add(a int, b int) int { // (a, b int) int {}
	return a + b
}

func get() (string, int, int) { // multiple values return
	return "golang", 12, 237
}

func ProcessIt1() func(a int) int {
	return func(b int) int {
		return 7
	}
}
func main() {
	fmt.Println("Functions")
	fmt.Println("Sum: ", add(3, 4))
	fmt.Println(get())
	lang1, lang2, _ := get() // _ replace in case of lang3 if we don't write the lang3.
	fmt.Println(lang1, lang2)

	fn := ProcessIt1()
	fmt.Println(fn(2))
}
