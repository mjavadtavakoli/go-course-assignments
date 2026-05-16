package main

import (
	"fmt"
	"time"
)

func mainn() {
	go Printnumbers()
	go Printletter()
	fmt.Println("done")
}

func Printletter() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Print(string(i), "")

		time.Sleep(time.Millisecond * 70)

	}

}
func Printnumbers() {
	for i := 1; i <= 50; i++ {
		fmt.Print(i)
		time.Sleep(time.Millisecond * 70)
	}

}
