package mockery

import (
	"github.com/golang/mock/gomock"
	"gitlab.betsys.com/go-guild-mocks/4-mockery/mocks"
	"testing"
)

func TestQuickMaths(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"Two plus two is four", 2, 2, 4},
		{"Minus one that's three, quick maths", 4, -1, 3},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			testObj := new(mocks.Math)

			qm := &QuickMaths{
				Math: testObj,
			}

			testObj.On("Sum", c.a, c.b).Return(c.want)
			// testObj.On("Sum", c.a, c.b).Return(func(a int, b int) int { return a + b })

			//testObj.AssertCalled(t, "Sum", c.a, c.b)
			//testObj.AssertNotCalled(t, "Sum", c.a, c.b)

			got := qm.Sum(c.a, c.b)

			testObj.AssertExpectations(t)

			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}
