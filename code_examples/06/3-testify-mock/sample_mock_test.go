package testify

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

/*
  Test objects
*/

// MyMockedService is a mocked object that implements an interface
// that describes an object that the code I am testing relies on.
type MyMockedService struct {
	mock.Mock
}

// DoSomething is a method on MyMockedService that implements some interface
// and just records the activity, and returns what the Mock object tells it to.
//
// In the real object, this method would do something useful, but since this
// is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.
func (m *MyMockedService) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

/*
  Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestTargetFuncThatDoesSomethingWithObj(t *testing.T) {

	// create an instance of our test object
	testObj := new(MyMockedService)

	// setup expectations
	testObj.On("DoSomething", 123).Return(true, nil)

	// call the code we are testing
	res, err := targetFuncThatDoesSomethingWithObj(testObj)

	// assert that the expectations were met
	testObj.AssertExpectations(t)

	assert.Equal(t, true, res, "they should be equal")
	assert.Nil(t, err)
}

func TestSomethingWithPlaceholder(t *testing.T) {

	// create an instance of our test object
	testObj := new(MyMockedService)

	// setup expectations with a placeholder in the argument list
	testObj.On("DoSomething", mock.Anything).Return(false, nil)

	// call the code we are testing
	res, err := targetFuncThatDoesSomethingWithObj(testObj)

	// assert that the expectations were met
	testObj.AssertExpectations(t)

	assert.Equal(t, false, res, "they should be equal")
	assert.Nil(t, err)
}
