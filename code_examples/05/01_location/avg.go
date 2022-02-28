package mypackage

type Average interface {
	Avg(float32) float32
}

type Number float32

func (n Number) Avg(avgNum float32) float32 {
	/*
		.
		. avg logic
		.
	*/

	return n.round()
}

func (n Number) round() float32 {
	/*
		.
		. round logic
		.
	*/

	return 0
}
