package main

import "fmt"

func mapp() {

	// =========================
	// 1) تعریف Map
	// =========================

	// روش 1
	ages := map[string]int{
		"Ali":  25,
		"Reza": 30,
	}
	fmt.Println("ages:", ages)

	// روش 2 با make
	scores := make(map[string]int)
	fmt.Println("scores:", scores)

	// =========================
	// 2) اضافه کردن / تغییر مقدار
	// =========================

	scores["math"] = 90
	scores["english"] = 85
	fmt.Println("scores after add:", scores)

	// تغییر مقدار
	scores["math"] = 95
	fmt.Println("scores after update:", scores)

	// =========================
	// 3) گرفتن مقدار
	// =========================

	mathScore := scores["math"]
	fmt.Println("math score:", mathScore)

	// =========================
	// 4) چک کردن وجود کلید
	// =========================

	val, exists := scores["science"]
	fmt.Println("exists:", exists, "value:", val)

	// =========================
	// 5) حذف کردن
	// =========================

	delete(scores, "english")
	fmt.Println("after delete:", scores)

	// =========================
	// 6) loop با range
	// =========================

	for key, value := range scores {
		fmt.Println(key, value)
	}

	// =========================
	// 7) گرفتن طول
	// =========================

	fmt.Println("len:", len(scores))

	// =========================
	// 8) map خالی یا nil
	// =========================

	var m1 map[string]int
	m2 := make(map[string]int)

	fmt.Println("m1 is nil:", m1 == nil)
	fmt.Println("m2 len:", len(m2))

	// =========================
	// 9) کپی map (دستی)
	// =========================

	src := map[string]int{"a": 1, "b": 2}
	dst := make(map[string]int)

	for k, v := range src {
		dst[k] = v
	}

	fmt.Println("dst:", dst)

}
