package main

import "fmt"

func main() {
	b := make([]int, 3, 5)
	fmt.Println("b:", b, "len:", len(b), "cap:", cap(b))

}

/*
len
cat
append
copy
*/
