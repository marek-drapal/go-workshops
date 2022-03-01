package testify

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

// require calls t.FailNow on failure (no return value)
func TestSomethingWithRequire(t *testing.T) {
	// assert equality
	require.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	require.NotEqual(t, 123, 456, "they should not be equal")

	type TestObject struct {
		Value string
	}

	var err error
	// assert for nil (good for errors)
	require.Nil(t, err)

	err = errors.New("something went wrong")

	// assert for not nil (good when you expect something)
	require.NotNil(t, err)

	// now we know that object isn't nil, we are safe to make
	// further assertions without causing any errors
	require.Equal(t, "something went wrong", err.Error())
}
