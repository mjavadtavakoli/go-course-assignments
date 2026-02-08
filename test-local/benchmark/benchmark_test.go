package benchmark

import "testing"

// تابعی که می‌خوایم سرعتش رو بسنجیم
func SumNumbers(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// Benchmark
func BenchmarkSumNumbers(b *testing.B) {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		SumNumbers(nums)
	}
}

//go test -bench=.
//go mod init benchmark
