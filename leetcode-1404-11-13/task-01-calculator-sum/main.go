package main

import (
    mohammadjavad "fmt"
)
func Sumnumber(numbers []int)int{

sum := 0
for _,num:= range numbers{
    sum += num
}
return sum
}

func main(){
    //mohammadjavad.Println("HI")
   mohammadjavad.Println("Test 1:", Sumnumber([]int{1, 2, 3, 4, 5})) // Output: 15
   mohammadjavad.Println("Test 2:", Sumnumber([]int{-1, 0, 1}))      // Output: 0
   mohammadjavad.Println("Test 3:", Sumnumber([]int{}))              // Output: 0
}

















/*package main

import "fmt"

func SumNumbers(numbers []int) int {
    sum := 0
    
    for _, num := range numbers {
        sum += num
    }
    
    return sum
}

func main() {
    // تست تابع با مثال‌های داده شده
    fmt.Println("Test 1:", SumNumbers([]int{1, 2, 3, 4, 5})) // Output: 15
    fmt.Println("Test 2:", SumNumbers([]int{-1, 0, 1}))      // Output: 0
    fmt.Println("Test 3:", SumNumbers([]int{}))              // Output: 0
}*/
