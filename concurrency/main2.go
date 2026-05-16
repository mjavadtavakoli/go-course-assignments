package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*func main() {
	var a, b, c int
	go func() {
		a = 1
	}()
	go func() {
		b = 2
	}()
	go func() {
		b = 3
	}()

	fmt.Println(a + b + c)
}
*/

//use sync.waitgroup(add-wait-doness)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		Printlettrs()
		wg.Done()
	}()

	go func() {
		Printnumbers()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("finished")
}

func Printlettrs() {

	for i := 'A'; i <= 'Z'; i++ {
		fmt.Print(string(i))
		//time.Sleep(time.Millisecond * 50)
	}

}

func Printnumbers() {
	for i := 1; i <= 50; i++ {
		fmt.Print(i)
		//time.Sleep(time.Millisecond * 50)
	}

}
