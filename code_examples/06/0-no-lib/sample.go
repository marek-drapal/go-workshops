package no_lib

type Math interface {
	Sum(a, b int) int
}

type QuickMaths struct {
	Math
}

func (qm *QuickMaths) Sum(a, b int) int {
	return qm.Math.Sum(a, b)
}

func (qm *QuickMaths) Minus(a, b int) int {
	return qm.Math.Sum(a, -b)
}
