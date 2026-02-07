package main

import "fmt"

func slice() {

	// =========================
	// 1) تعریف Slice
	// =========================

	// روش 1: تعریف با مقدار اولیه
	a := []int{1, 2, 3}
	fmt.Println("a:", a)

	// روش 2: با make (len=3 , cap=5)
	b := make([]int, 3, 5)
	fmt.Println("b:", b, "len:", len(b), "cap:", cap(b))

	// روش 3: nil slice
	var c []int
	fmt.Println("c:", c, "len:", len(c))

	// =========================
	// 2) مقداردهی
	// =========================

	b[0] = 10
	b[1] = 20
	b[2] = 30
	fmt.Println("b after set:", b)

	// =========================
	// 3) append (اضافه کردن)
	// =========================

	b = append(b, 40)
	b = append(b, 50, 60)
	fmt.Println("b after append:", b)

	// =========================
	// 4) loop روی slice
	// =========================

	for i, v := range b {
		fmt.Println("index:", i, "value:", v)
	}

	// =========================
	// 5) برش (Slicing)
	// =========================

	x := []int{1, 2, 3, 4, 5}
	part := x[1:4] // از اندیس 1 تا قبل 4
	fmt.Println("part:", part)

	// =========================
	// 6) حذف عنصر از slice
	// =========================

	// حذف اندیس 2
	x = append(x[:2], x[3:]...)
	fmt.Println("x after delete:", x)

	// =========================
	// 7) کپی slice
	// =========================

	src := []int{7, 8, 9}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("dst:", dst)

	// =========================
	// 8) مقایسه nil یا خالی
	// =========================

	var n []int
	e := []int{}

	fmt.Println(n == nil) // true
	fmt.Println(len(e) == 0)

	// =========================
	// 9) append یک slice به slice دیگر
	// =========================

	a1 := []int{1, 2}
	a2 := []int{3, 4}

	a1 = append(a1, a2...)
	fmt.Println("merged:", a1)

	// =========================
	// 10) capacity growth
	// =========================

	nums := make([]int, 0)
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
		fmt.Println("len:", len(nums), "cap:", cap(nums))
	}

}
