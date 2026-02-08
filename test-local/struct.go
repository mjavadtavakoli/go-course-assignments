package main

import "fmt"

// =========================
// 1) تعریف Struct
// =========================

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func Structt() {

	// =========================
	// 2) ساختن Struct
	// =========================

	u1 := User{
		ID:    1,
		Name:  "Ali",
		Email: "ali@gmail.com",
		Age:   25,
	}

	fmt.Println(u1)

	// =========================
	// 3) دسترسی به فیلدها
	// =========================

	fmt.Println(u1.Name)
	fmt.Println(u1.Age)

	// =========================
	// 4) تغییر مقدار فیلد
	// =========================

	u1.Age = 26
	fmt.Println(u1.Age)

	// =========================
	// 5) Struct بدون نام فیلد
	// (ترتیب مهم است)
	// =========================

	u2 := User{2, "Reza", "reza@gmail.com", 30}
	fmt.Println(u2)

	// =========================
	// 6) Pointer به Struct
	// =========================

	u3 := &User{
		ID:   3,
		Name: "Sara",
		Age:  22,
	}

	fmt.Println(u3.Name) // خودکار Dereference

	// =========================
	// 7) Struct داخل Struct
	// =========================

	type Address struct {
		City string
		Zip  string
	}

	type Customer struct {
		Name    string
		Address Address
	}

	c := Customer{
		Name: "Mohammad",
		Address: Address{
			City: "Tehran",
			Zip:  "12345",
		},
	}

	fmt.Println(c.Address.City)

	// =========================
	// 8) Slice از Struct
	// =========================

	users := []User{
		u1,
		u2,
	}

	fmt.Println(users)

	// =========================
	// 9) Loop روی Struct ها
	// =========================

	for _, user := range users {
		fmt.Println(user.Name, user.Age)
	}

}
