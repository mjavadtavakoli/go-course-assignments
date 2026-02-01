/*Fix this code:

package main

import "fmt"

func main() {
    var quantity int = 5
    var unitPrice float64 = 29.99
    
    // ❌ This won't compile
    total := quantity * unitPrice
    
    fmt.Printf("Total: $%.2f\n", total)
}*/

package main

import "fmt"

func main() {
    var quantity int = 5
    var unitPrice float64 = 29.99
    
    // ✅ Convert quantity to float64
    total := float64(quantity) * unitPrice
    
    fmt.Printf("Total: $%.2f\n", total)
}