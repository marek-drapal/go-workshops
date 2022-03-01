package no_lib

import "testing"

// mock structure
type mockMath struct {
	Calls int
}

func (m *mockMath) Sum(a, b int) int {
	m.Calls++
	return a + b
}

// actual test
func TestQuickMaths(t *testing.T) {
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
			mock := new(mockMath)

			qm := &QuickMaths{
				Math: mock,
			}

			got := qm.Sum(c.a, c.b)

			if mock.Calls != 1 {
				t.Errorf("Expected 1 call to Math.Sum, got %d", mock.Calls)
			}

			if got != c.want {
				t.Errorf("got %d, want %d", got, c.want)
			}
		})
	}
}
