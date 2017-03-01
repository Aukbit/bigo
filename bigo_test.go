package bigo

import (
	"testing"

	"github.com/paulormart/assert"
)

var slice10 = make([]int, 10, 10)
var slice1000 = make([]int, 1000, 1000)
var slice1M = make([]int, 1000000, 1000000)

// --
func TestIsFirstZero(t *testing.T) {
	assert.Equal(t, true, IsFirstZero(slice10))
	assert.Equal(t, true, IsFirstZero(slice1M))
}

func BenchmarkIsFirstZero10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsFirstZero(slice10)
	}
}

func BenchmarkIsFirstZero1M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsFirstZero(slice1M)
	}
}

// --
func TestContains(t *testing.T) {
	slice10[9] = 1
	slice1M[999999] = 1
	assert.Equal(t, true, Contains(slice10, 1))
	assert.Equal(t, true, Contains(slice1M, 1))
}

func BenchmarkContains10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Contains(slice10, 1)
	}
}

func BenchmarkContains1M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Contains(slice1M, 1)
	}
}

// --
func TestContainsDuplicates(t *testing.T) {
	slice10[9] = 1
	slice10[8] = 1
	slice1000[999] = 1
	slice1000[998] = 1
	assert.Equal(t, true, ContainsDuplicates(slice10))
	assert.Equal(t, true, ContainsDuplicates(slice1000))
}

func BenchmarkContainsDuplicates10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainsDuplicates(slice10)
	}
}

func BenchmarkContainsDuplicates1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainsDuplicates(slice1000)
	}
}

// --

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 55, Fibonacci(10))
}

func BenchmarkFibonacci10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}
func BenchmarkFibonacci20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}
