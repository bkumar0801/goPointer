package single_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"thoughtworks.com/goPointer/single"
)

func TestShouldGiveValueBack(t *testing.T)  {
	value := single.GetValueBack(1)
	assert.EqualValues(t, value, 1)
	value = single.GetValueBack(1000)
	assert.EqualValues(t, value, 1000)
	value = single.GetValueBack('a')
	assert.EqualValues(t, value, 'a')
}

func TestShouldGetOriginalValue(t *testing.T)  {
	var myValue int = 23   //TODO: check with other data type, should throw error
	value  := single.GetOriginalValue(&myValue)
	assert.EqualValues(t, value, myValue)
}

func TestFunQuiz(t *testing.T)  {
	var r int = 10
	var p *int = &r
	single.FunQuiz(p)
	assert.EqualValues(t, *p, r)
	single.FunQuiz(&r)
	assert.EqualValues(t, *p, r)
}

func TestGetAddress(t *testing.T) {
	p := single.GetValueRef(10);
	*p = 30
	assert.EqualValues(t, *p, 30)
}
