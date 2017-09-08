package multiple_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"thoughtworks.com/goPointer/multiple"
)

func TestShouldGiveMyValueBack(t *testing.T)  {
	value := multiple.GetMyValueBack(100)
	assert.EqualValues(t, value, 100)
}
