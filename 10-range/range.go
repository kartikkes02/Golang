package main
import "fmt"

// range -> used for the ierating over data structures.
func main() {
	fmt.Println("Range")

	// create slices
	nums := []int{6, 7, 8}

	fmt.Println("Index ", "Num")
	sum := 0
	for i, num := range nums {
		fmt.Println(i, "     ", num)
		sum += num
	}
	fmt.Println("Sum: ", sum)

	fmt.Println("For Maps: ")
	maps := make(map[string]int)
	maps["name"] = 4
	maps["address"] = 6

	fmt.Println("Index ", "Numss")
	for l, numss := range maps {
		fmt.Println(l, " ", numss)
	}

	// unicode code point rune: "golang"
	// starting point of rune: k
	for k, b := range "golang" {
		fmt.Println(k, b)
	}
}
