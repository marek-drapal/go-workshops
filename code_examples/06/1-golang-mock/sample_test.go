package golang_mock

import (
	"github.com/golang/mock/gomock"
	"gitlab.betsys.com/go-guild-mocks/1-golang-mock/mocks"
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
			m := mocks.NewMockMath(ctrl)

			qm := &QuickMaths{
				Math: m,
			}

			m.EXPECT().
				Sum(gomock.Eq(c.a), gomock.Eq(c.b)).
				Return(c.want)
			// AnyTimes() - don't care how many times this method is called (or not called)

			got := qm.Sum(c.a, c.b)
			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}
