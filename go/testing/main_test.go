// testing and benchmarking
package main

import (
	// "fmt"
	"math/rand"
	"testing"
)

func GenRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func SumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func TestGenRandomSlice(t *testing.T) {
	size := 100
	slice := GenRandomSlice(size)
	if len(slice) != size {
		t.Errorf("expected slice size: %d; received: %d", size, len(slice))
	}
}

func BenchmarkGenRandomSlice(b *testing.B) {
	for b.Loop() {
		GenRandomSlice(100)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := GenRandomSlice(100)
	b.ResetTimer()
	for b.Loop() {
		SumSlice(slice)
	}
}

// func Add(a, b int) int {
// 	return a + b
// }

// -----------------BENCHMARKING----------------
// func BenchmarkAdd(b *testing.B) {
// 	for b.Loop() {
// 		Add(2, 3)
// 	}
// }
//
// func BenchmarkAddMedium(b *testing.B) {
// 	for b.Loop() {
// 		Add(200, 300)
// 	}
// }
//
// func BenchmarkAddLarge(b *testing.B) {
// 	for b.Loop() {
// 		Add(2000, 3000)
// 	}
// }
//
// // ----------------TESTING---------------------
// func TestAddSubtests(t *testing.T) {
// 	tests := []struct{ a, b, expected int }{
// 		{2, 3, 5},
// 		{0, 0, 0},
// 		{-2, 2, 0},
// 	}
//
// 	for _, test := range tests {
// 		t.Run(fmt.Sprintf("Add(%d, %d)", test.a, test.b), func(t *testing.T) {
// 			result := Add(test.a, test.b)
// 			if result != test.expected {
// 				t.Errorf("result = %d; want: %d", result, test.expected)
// 			}
// 		})
// 	}
// }
//
// func TestAddTableDriven(t *testing.T) {
// 	tests := []struct{ a, b, expected int }{
// 		{2, 3, 5},
// 		{1, 2, 3},
// 		{-1, 1, 0},
// 	}
//
// 	for _, test := range tests {
// 		result := Add(test.a, test.b)
// 		if result != test.expected {
// 			t.Errorf("Add(%d, %d) = %d; want: %d", test.a, test.b, result, test.expected)
// 		}
// 	}
// }
//
// func TestAdd(t *testing.T) {
// 	result := Add(2, 3)
// 	expected := 5
// 	if result != expected {
// 		t.Errorf("Add(2, 3) = %d; want: %d", result, expected)
// 	}
// }
