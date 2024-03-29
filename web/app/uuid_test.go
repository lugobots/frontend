package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniquer_New_WillReturnErrorWhenListRunsOut(t *testing.T) {
	U := Uniquer{}
	list = []string{"a"}
	s, err := U.New()
	assert.Nil(t, err)
	assert.Equal(t, "a", s)

	s, err = U.New()
	assert.Equal(t, ErrNoMoreUUIDs, err)
	assert.Equal(t, "", s)
}

func TestUniquer_New_ReturnPseudoRandom(t *testing.T) {

	list = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	A := Uniquer{}
	sa, errA := A.New()

	B := Uniquer{}
	sb, errB := B.New()

	assert.Nil(t, errA)
	assert.Nil(t, errB)
	assert.NotEqual(t, sa, sb)
}

func TestUniquer_New_ReturnUniqueValue(t *testing.T) {

	list = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	returnedValues := map[string]bool{}

	A := Uniquer{}
	var err error
	for {
		var s string
		s, err = A.New()
		if err != nil {
			break
		}
		_, alreadyReturned := returnedValues[s]
		assert.False(t, alreadyReturned)
		returnedValues[s] = true
	}
	assert.Equal(t, ErrNoMoreUUIDs, err)
	assert.Equal(t, len(list), len(returnedValues))
}
