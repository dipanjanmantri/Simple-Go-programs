package fib

import "testing"
import "fmt"

func check(t *testing.T, k int, expectation int) {
	r := fibonacci(k)
	if r != expectation{
		t.Error("Expected ", expectation, " got ", r)
	}
}

func TestFib(t1 *testing.T) {
	check(t1, 0, 0)
	check(t1, 1, 1)
	check(t1, 2, 1)
	check(t1, 3, 2)
	check(t1, 4, 3)
	check(t1, 5, 5)
	check(t1, 10, 55)

	fmt.Println(fibonacci(20))
}
