package arr_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"thoughtworks.com/goPointer/arr"
)

func TestShouldGiveMyValueBack(t *testing.T)  {
	fib := [6]int{0, 1, 1, 2, 3, 5}
	arr.DoubleFib(&fib)
	assert.EqualValues(t, fib, [6]int{0, 2, 2, 4, 6, 10})
}

func TestShouldManipulateStr(t *testing.T)  {
	fib := [6]string{"0", "1", "1", "2", "3", "5"}
	arr.StringManip(&fib)
	assert.EqualValues(t, fib, [6]string{"0", "1", "1", "2", "3", "5"})
}
