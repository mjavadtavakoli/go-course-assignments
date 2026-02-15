// توی این فایل ها هم فانکشن های تست و هم فانکشن های بنچ مارک نوشته میشود
package main

import (
	"testing"
)

type TestCase struct {
	a              int
	b              int
	expectedresult int
}

func TestSum(t *testing.T) {
	testCases := []TestCase{
		{a: 1, b: 2, expectedresult: 3},
		{a: -1, b: 1, expectedresult: 0},
		{a: 0, b: 0, expectedresult: 0},
		{a: 10, b: 20, expectedresult: 30},
	}
	for _, testCase := range testCases {
		result := Sum(testCase.a, testCase.b)
		if result != testCase.expectedresult {
			t.Errorf("Sum(%d, %d) = %d, expected %d", testCase.a, testCase.b, result, testCase.expectedresult)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}
