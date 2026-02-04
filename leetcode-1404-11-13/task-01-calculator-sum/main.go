package main

import (
	mohammadjavad "fmt"
)

func Sumnumber(numbers []int) int {

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	//mohammadjavad.Println("HI")
	mohammadjavad.Println("Test 1:", Sumnumber([]int{1, 2, 3, 4, 5}))   // Output: 15
	mohammadjavad.Println("Test 2:", Sumnumber([]int{-1, 0, 1}))        // Output: 0
	mohammadjavad.Println("Test 3:", Sumnumber([]int{}))                // Output: 0
	mohammadjavad.Println("Test 1:", Sumnumber([]int{21, 4, 15, 0, 1})) // Output: 15

}

/* guide code

first : write sumnumber function with slice
then : create variable
after that : write loop with range
*/
