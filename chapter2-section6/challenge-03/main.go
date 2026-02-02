package main

import (
    "fmt"
    "strings"
)

func isPalindrome(x int) bool {
    if x < 0 || (x%10 == 0 && x != 0) {
        return false
    }
    
    reversedHalf := 0
    for x > reversedHalf {
        reversedHalf = reversedHalf*10 + x%10
        x /= 10
    }
    
    return x == reversedHalf || x == reversedHalf/10
}

func main() {
    tests := []struct {
        input    int
        expected bool
    }{
        {121, true},     // Palindrome (odd digits)
        {1221, true},    // Palindrome (even digits)
        {-121, false},   // Negative number
        {10, false},     // Ends with 0
        {0, true},       // Zero
        {12345, false},  // Non-palindrome
        {12321, true},   // Large palindrome
        {7, true},       // Single digit
        {1001, true},    // Palindrome with zeros
    }
    
    // Display header
    separator := strings.Repeat("═", 50)
    fmt.Println("🧠 Palindrome Number Validator")
    fmt.Println(separator)
    
    // Run tests
    allPassed := true
    for _, tc := range tests {
        result := isPalindrome(tc.input)
        passed := result == tc.expected
        
        if !passed {
            allPassed = false
        }
        
        // Colored output (optional)
        if passed {
            fmt.Printf("✅ ")
        } else {
            fmt.Printf("❌ ")
        }
        
        fmt.Printf("isPalindrome(%6d) = %-5v | Expected: %-5v\n",
            tc.input, result, tc.expected)
    }
    
    // Summary
    fmt.Println(separator)
    if allPassed {
        fmt.Println("🎉 All tests passed!")
    } else {
        fmt.Println("⚠️  Some tests failed!")
    }
    fmt.Println(separator)
    
    // Additional examples
    fmt.Println("\n📋 Quick Examples:")
    examples := []int{121, -121, 10, 0, 12321, 12345}
    for _, num := range examples {
        fmt.Printf("  isPalindrome(%d) = %v\n", num, isPalindrome(num))
    }
}