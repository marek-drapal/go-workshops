package testify

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	var err error
	// assert for nil (good for errors)
	assert.Nil(t, err)

	err = errors.New("something went wrong")
	// assert for not nil (good when you expect something)
	if assert.NotNil(t, err) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "something went wrong", err.Error())
	}
}

func TestWithNew(t *testing.T) {
	// you can create a new instance of the assertions (no need to repeat t)
	assert := assert.New(t)

	// assert equality
	assert.Equal(123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(123, 456, "they should not be equal")

	var err error
	// assert for nil (good for errors)
	assert.Nil(err)

	err = errors.New("something went wrong")
	// assert for not nil (good when you expect something)
	if assert.NotNil(err) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal("something went wrong", err.Error())
	}
}
