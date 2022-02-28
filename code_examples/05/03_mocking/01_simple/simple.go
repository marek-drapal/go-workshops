package simple

type Plain interface {
	TakeOff() bool
}

type Booing struct{}

func (b *Booing) TakeOff() bool {
	/*
		.
		. some logic
		.
	*/

	return true
}
