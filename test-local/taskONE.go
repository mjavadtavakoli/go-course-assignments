/*We want to create a program that:

has a list of student names

can add a new student

and finally prints all the students
*/

package main

import "fmt"

func main() {

	// لیست اولیه دانشجوها
	students := []string{"Ali", "Reza", "Sara"}

	// اضافه کردن دانشجوی جدید
	students = append(students, "Mohammad")

	// چاپ همه دانشجوها
	for i, name := range students {
		fmt.Println(i+1, name)
	}
}
