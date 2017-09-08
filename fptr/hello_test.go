package fptr_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"thoughtworks.com/goPointer/fptr"
)

func TestShouldGiveMyValueBack(t *testing.T)  {
	say := fptr.SayHello
	greetings := say("World!")
	assert.EqualValues(t, greetings, "Hello,World!")
}

func TestShouldCalculate(t *testing.T) {
	value := fptr.Calculate(fptr.Plus, 1,2)
	assert.EqualValues(t, value, 3)
	value = fptr.Calculate(fptr.Multiply, 4,90)
	assert.EqualValues(t, value, 360)
	value = fptr.Calculate(fptr.Minus, 10,2)
	assert.EqualValues(t, value, 8)
}

