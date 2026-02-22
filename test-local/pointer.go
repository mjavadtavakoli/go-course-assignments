package main

import "fmt"

func mains() {
	poin()

}

func poin() {
	x := 10
	p := &x
	fmt.Println(*p)
}
