package bigo

import (
	"fmt"
	"strconv"
)

// IsFirstZero O(1) describes an algorithm that the time and space of it's execution
// is always the same regardless the size of the input dataset
func IsFirstZero(n []int) bool {
	return n[0] == 0
}

// Contains O(N) describes an algorithm that the time and space of it's execution
// will grow linear with the size of the input dataset
func Contains(n []int, value int) bool {
	for _, v := range n {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsDuplicates O(N^2) describes an algorithm that the time and space of it's execution
// is directly proportional of the size of the input dataset
func ContainsDuplicates(n []int) bool {
	for i, a := range n {
		for j, b := range n {
			if i != j && a == b && a != 0 {
				return true
			}
		}
	}
	return false
}

// Fibonacci O(2^N) describes an algorithm that the time and space of it's execution
// doubles on each iteration increasing exponentially
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// BinarySearch O(log N) describes an algorithm that the time and space of it's execution
// halves on each iteration
// package sort.Search implements binary search

// FizzBuzz prints the numbers from 1 to 100.
// But for multiples of three print “Fizz” instead of the number and for the
// multiples of five print “Buzz”.
// For numbers which are multiples of both three and five print “FizzBuzz”."
func Fizzbuzz() {

	for i := 1; i <= 100; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("buzz")
			continue
		}
		fmt.Println(i)
	}
}

func FizzbuzzRecursive(n int, out string) string {
	fmt.Println(out)
	if n > 100 {
		return ""
	}
	if n%5 == 0 && n%3 == 0 {
		return FizzbuzzRecursive(n+1, "fizzbuzz")
	}
	if n%3 == 0 {
		return FizzbuzzRecursive(n+1, "fizz")
	}
	if n%5 == 0 {
		return FizzbuzzRecursive(n+1, "buzz")
	}
	return FizzbuzzRecursive(n+1, strconv.Itoa(n))
}
