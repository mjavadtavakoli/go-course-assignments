package main

import "fmt"

var target = 7
var input int

func main() {

	fmt.Sscanf(&input)

	if input > target {
		fmt.Println("Too big")
	}else if input < target {
		fmt.Println("Too small")
	}else input == target {
		fmt.Println("curecct")
	}
}
	

