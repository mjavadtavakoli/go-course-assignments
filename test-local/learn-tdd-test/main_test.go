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

type TestcaseManfi struct {
	f              int
	h              int
	expectedresult int
}

func TestManfi(t *testing.T) {
	testcasemanfi := []TestcaseManfi{
		{f: 1, h: 2, expectedresult: 3},
		{f: -1, h: 1, expectedresult: 0},
		{f: 0, h: 0, expectedresult: 0},
		{f: 10, h: 20, expectedresult: 30},
	}
	for _, testCase := range testcasemanfi {
		result := Sum(testCase.f, testCase.h)
		if result != testCase.expectedresult {
			t.Errorf("Sum(%d, %d) = %d, expected %d", testCase.f, testCase.h, result, testCase.expectedresult)
		}
	}
}
