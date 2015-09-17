package question1

import "testing"

func Test_fib(t *testing.T) {
    var result1 int
    var result2 int
    var input1 int = 5
    var input2 int = 4

    result1 = fib(input1)
    if result1!=5 {
        t.Error("Expected 5, got ",result1)
    }
    result2 = fib(input2)
    if result2!=3 {
        t.Error("Expected 3, got ",result2)
    }
}