package main

import (
	"fmt"
	"unicode/utf8"
)

func rbs() {

	// =========================
	// 1) String
	// =========================

	text := "Hello Go"
	fmt.Println(text)

	fmt.Println("len(text):", len(text)) // تعداد بایت

	// =========================
	// 2) Byte (uint8)
	// =========================

	b := text[0]
	fmt.Println("first byte:", b)
	fmt.Println("first byte as char:", string(b))

	// =========================
	// 3) Rune (int32)
	// =========================

	runes := []rune(text)
	fmt.Println("runes:", runes)
	fmt.Println("first rune:", runes[0])
	fmt.Println("first rune as char:", string(runes[0]))

	// =========================
	// 4) Unicode Example
	// =========================

	persian := "سلام"
	fmt.Println(persian)

	fmt.Println("len(persian):", len(persian)) // بایت
	fmt.Println("rune count:", utf8.RuneCountInString(persian))

	// =========================
	// 5) Loop روی String با range
	// =========================

	for i, r := range persian {
		fmt.Println(i, string(r))
	}

	// =========================
	// 6) تبدیل‌ها
	// =========================

	s := "ABC"

	bytes := []byte(s)
	runes2 := []rune(s)

	fmt.Println(bytes)
	fmt.Println(runes2)

	fmt.Println(string(bytes))
	fmt.Println(string(runes2))

}
