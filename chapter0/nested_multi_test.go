package chapter0_test

import (
	"testing"
	"github.com/xiaochuan-dev/numerical-analysis/chapter0"
)

func TestNestedMulti(t *testing.T)  {
	a := []int{2, 3, -3, 5, -1}
	var x int = 2
	
	res := chapter0.NestedMulti(a, x)

	if res != 53 {
		t.Fatal("TestNestedMulti bad answer")
	}
}
