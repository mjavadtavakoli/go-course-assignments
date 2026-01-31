// (Before) :

/*package main
import "fmt"
func main(){fmt.Println("Hello, Go!") if true{fmt.Println("Formatting is bad")}}*/






// (after) : 
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")

	if true {
		fmt.Println("Formatting is great")
	}
}

